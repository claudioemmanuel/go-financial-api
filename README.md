# Financial Control Backend

This repository contains the backend for the Financial Control project, which aims to help individuals manage their personal finances. The backend is built using Golang, Gin web framework, and Gorm ORM, following the hexagonal architecture pattern.

## Features

- Account Management: Add and manage multiple accounts
- Transaction Tracking: Record income and expenses
- Budgeting and Planning: Set spending targets and monitor progress
- Financial Reports: Generate summaries and visualizations of financial data
- Alerts and Notifications: Send reminders for future payments or budget limits

## Requirements

- Go 1.19
- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
    
```bash
git clone https://github.com/claudioemmanuel/financial-api.git
cd financial-api
```

2. Set up environment variables:
Create a `.env` file in the project root and add the required environment variables:
    
```bash
cp .env.example .env
```

3. Build and run the Docker containers:

```bash
docker-compose up --build
```

The API should now be running on `http://localhost:8080`.
## Running Tests

1. Make sure you have Go and SQLite installed on your system.

2. Run the tests:

```bash
go test -v ./...
```

## API Documentation

The API is documented using Swagger, an open-source framework for API documentation. To view the documentation, follow these steps:

1. Generate the Swagger documentation:
   - On your local machine, run `swag init`. If you don't have Swag CLI installed, follow the [Swag CLI installation instructions](https://github.com/swaggo/swag#installation).
   - Alternatively, you can generate the documentation inside the container by running `docker-compose run swagger`.

2. Once the documentation is generated, it will be available in the `docs` folder. Start the API by running `docker-compose up -d` if it's not already running.

3. Open your browser and navigate to `http://localhost:8080/swagger/index.html`. You should see the Swagger UI displaying the API documentation.

Note: When working on your local environment, you will need to generate the Swagger documentation manually whenever you make changes to your route handlers that require updating the documentation.

## License

This project is licensed under the MIT License - see the [MIT LICENSE](https://opensource.org/licenses/MIT) for details.

## Acknowledgments

- [Golang](https://golang.org/)
- [Gin](https://gin-gonic.com)
- [Gorm](https://gorm.io/)
- [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Swagger](https://swagger.io/)

