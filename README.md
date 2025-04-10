# Go Boilerplate

A robust and scalable boilerplate for building API services in Go.

## Features

- **Modular Architecture**: Organized project structure promoting code reusability and separation of concerns.
- **Environment Configuration**: Utilizes `.env` files for managing environment-specific settings.
- **Database Integration**: Supports database migrations and interactions.
- **Docker Support**: Includes Docker configurations for containerized deployment.
- **Makefile Commands**: Simplifies common tasks such as building, running, and testing the application.

## Project Structure

```plaintext
├── ci/                  # Continuous integration configurations
├── cmd/                 # Application entry points
├── internal/            # Private application and library code
├── migration/           # Database migration files
├── .env.example         # Example environment variables file
├── .gitignore           # Git ignore file
├── Dockerfile           # Docker configuration
├── LICENSE              # License information
├── Makefile             # Makefile for task automation
├── docker-compose.yml   # Docker Compose configuration
├── go.mod               # Go module file
├── go.sum               # Go module dependencies
└── main.go              # Main application entry point
```

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Docker (for containerization)
- Make (for task automation)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/nibak/go-boilerplate.git
   cd go-boilerplate
2. **Set up environment variables**:
    Rename .env.example to .env and configure the necessary environment variables.

3. **Run the application**:
    ```bash
    make run
    ```
# Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

# License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE.md) file for details.