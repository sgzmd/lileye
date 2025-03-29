# Android Notification Backend

A lightweight backend service for storing and viewing Android notifications.

## Project Structure

```
.
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   ├── models/          # Data models and database schemas
│   ├── handlers/        # HTTP request handlers
│   ├── services/        # Business logic
│   └── storage/         # Database operations
├── web/
│   ├── static/          # Static assets (CSS, JS)
│   └── templates/       # HTML templates
├── tests/               # Test files
└── docs/               # Documentation
```

## Prerequisites

- Go 1.21 or later
- SQLite3

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/lileye/backend.git
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run tests:
   ```bash
   go test ./...
   ```

4. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

## Development

- All new features should be developed on separate branches
- Write tests before implementing features
- Run tests before committing changes
- Update README.md with new changes

## Testing

The project uses Go's built-in testing framework. Tests can be run with:
```bash
go test ./...
```

For verbose output:
```bash
go test -v ./...
```

## API Documentation

[API documentation will be added as endpoints are implemented]

## Frontend

The frontend is built using:
- Alpine.js for interactivity
- Tailwind CSS for styling
- Native JavaScript for date handling

[Frontend documentation will be added as components are implemented] 