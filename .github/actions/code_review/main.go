package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"

	"errors"

	"github.com/sashabaranov/go-openai"
)

type CodeReviewRequest struct {
	CodeSnippet string
}

type CodeReviewResponse struct {
	Comments []string
}

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

	// Create the CodeReviewRequest
	codeReviewRequest := CodeReviewRequest{}

	// Perform the code review using GPT API
	_, err = ReviewCode(ctx, client, prEvent, codeReviewRequest)
	if err != nil {
			fmt.Printf("Failed to review code: %v\n", err)
			os.Exit(1)
	}
}

func ReviewCode(ctx context.Context, client *github.Client, event *github.PullRequestEvent, request CodeReviewRequest) (*CodeReviewResponse, error) {

    // Get the list of changed files in the PR
    files, _, err := client.PullRequests.ListFiles(ctx, event.Repo.Owner.GetLogin(), event.Repo.GetName(), event.GetNumber(), nil)
    if err != nil {
        return nil, fmt.Errorf("error getting changed files: %w", err)
    }

    // Generate a prompt for the GPT API
    prompt := "Review the following code changes:\n"
    for _, file := range files {
        prompt += fmt.Sprintf("File: %s\nPatch:\n%s\n", file.GetFilename(), file.GetPatch())
    }

    // Call the GPT API using the go-openai package
    review, err := ChatGPTReview(ctx, prompt)
    if err != nil {
        return nil, fmt.Errorf("error getting GPT review: %w", err)
    }

    // Post the GPT-generated review as a comment on the pull request
    comment := &github.IssueComment{
        Body: github.String(review),
    }
    _, _, err = client.Issues.CreateComment(ctx, event.Repo.Owner.GetLogin(), event.Repo.GetName(), event.GetNumber(), comment)
    if err != nil {
        return nil, fmt.Errorf("error posting review comment: %w", err)
    }

    response := &CodeReviewResponse{
        Comments: []string{review},
    }

    return response, nil
}

func ChatGPTReview(ctx context.Context, prompt string) (string, error) {

	// Get the OpenAI API key from the environment variable
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
			return "", errors.New("OPENAI_API_KEY is not set")
	}

	// Initialize the OpenAI client
	client := openai.NewClient(openaiAPIKey)

	// Create a completion prompt for code review
	completionRequest := openai.CompletionRequest{
			Model:  "text-davinci-002",
			Prompt: prompt,
			MaxTokens: 150,
			N: 1,
			Stop: []string{"\n"},
			Temperature: 0.5,
	}

	// Request completion from the OpenAI API
	completions, err := client.CreateCompletion(ctx, completionRequest)
	if err != nil {
			return "", err
	}

	// Check if any completions are returned
	if len(completions.Choices) == 0 {
			return "", errors.New("no completions returned")
	}

	// Get the first completion text
	result := completions.Choices[0].Text

	return result, nil
}