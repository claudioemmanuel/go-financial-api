package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"

	"errors"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// Get the token from the environment variable
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("GITHUB_TOKEN is not set")
		os.Exit(1)
	}

	// Authenticate with the token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Get the pull request event data
	eventData := os.Getenv("GITHUB_EVENT_DATA")
	if eventData == "" {
		fmt.Println("GITHUB_EVENT_DATA is not set")
		os.Exit(1)
	}

	event, err := github.ParseWebHook("pull_request", []byte(eventData))
	if err != nil {
		fmt.Printf("Failed to parse webhook: %v\n", err)
		os.Exit(1)
	}

	// Perform a type assertion to convert the event to a pull request event
	prEvent, ok := event.(*github.PullRequestEvent)
	if !ok {
		fmt.Println("Failed to convert event to pull request event")
		os.Exit(1)
	}

	// Perform the code review using GPT API
	err = ReviewCode(ctx, client, prEvent)
	if err != nil {
		fmt.Printf("Failed to review code: %v\n", err)
		os.Exit(1)
	}
}

func ReviewCode(ctx context.Context, client *github.Client, event *github.PullRequestEvent) error {
	// Get the list of changed files in the PR
	files, _, err := client.PullRequests.ListFiles(ctx, event.Repo.Owner.GetLogin(), event.Repo.GetName(), event.GetNumber(), nil)
	if err != nil {
		return fmt.Errorf("error getting changed files: %w", err)
	}

	// Get the raw diff for the PR
	diff, _, err := client.PullRequests.GetRaw(ctx, event.Repo.Owner.GetLogin(), event.Repo.GetName(), event.GetNumber(), github.RawOptions{Type: github.Diff})
	if err != nil {
		return fmt.Errorf("error getting raw diff: %w", err)
	}

	for _, file := range files {
		// Generate a prompt for the GPT API
		prompt := fmt.Sprintf("Please review the following code changes and provide feedback on the code quality, design decisions, and potential improvements:\nFile: %s\n", file.GetFilename())

		fileDiff := extractFileDiff(file.GetFilename(), diff)
		prompt += fmt.Sprintf("Diff:\n%s\n", fileDiff)

		// Call the GPT API using the go-openai package
		review, err := chatGPTReview(ctx, prompt)
		if err != nil {
			return fmt.Errorf("error getting GPT review: %w", err)
		}

		// Check 422 Validation Failed [{Resource:IssueComment Field:data Code:unprocessable Message:Body cannot be blank}]
		if strings.TrimSpace(review) == "" {
			review = "No review provided"
		}

		result := fmt.Sprintf("ChatGPT's response about `%s`:\n %s", file.GetFilename(), review)

		// Post the GPT-generated review as a comment on the pull request
		comment := &github.IssueComment{
			Body: github.String(result),
		}

		_, _, err = client.Issues.CreateComment(ctx, event.Repo.Owner.GetLogin(), event.Repo.GetName(), event.GetNumber(), comment)
		if err != nil {
			return fmt.Errorf("error posting review comment: %w", err)
		}
	}

	return nil
}

func chatGPTReview(ctx context.Context, prompt string) (string, error) {

	// Get the OpenAI API key from the environment variable
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		return "", errors.New("OPENAI_API_KEY is not set")
	}

	// Initialize the OpenAI client
	client := openai.NewClient(openaiAPIKey)

	// Create a completion prompt for code review
	completionRequest := openai.CompletionRequest{
		Model:       "text-davinci-002",
		Prompt:      prompt,
		MaxTokens:   150,
		N:           1,
		Stop:        []string{"\n"},
		Temperature: 0.5,
	}

	// Request completion from the OpenAI API
	completions, err := client.CreateCompletion(ctx, completionRequest)
	if err != nil {
		return "", err
	}

	fmt.Println(completions)

	// Check if any completions are returned
	if len(completions.Choices) == 0 {
		return "", errors.New("no completions returned")
	}

	// Get the first completion text
	result := completions.Choices[0].Text

	return result, nil
}

func extractFileDiff(fileName, diff string) string {
	fileDiff := ""
	lines := strings.Split(diff, "\n")
	inTargetFile := false

	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git") {
			inTargetFile = false
		}

		if strings.HasPrefix(line, "+++ b/"+fileName) {
			inTargetFile = true
		}

		if inTargetFile {
			fileDiff += line + "\n"
		}
	}

	return fileDiff
}
