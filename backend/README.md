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

### Endpoints

#### POST /api/notifications
Create a new notification.

Request body:
```json
{
    "title": "Notification Title",
    "message": "Notification Message",
    "timestamp": "2021-01-01T00:00:00Z",
    "package_name": "com.example.app",
    "from": "Other User Name",
    "device_id": "abc1234"
}
```

#### GET /api/notifications/:id
Get a notification by ID.

#### GET /api/notifications/device/:deviceID
Get all notifications for a specific device.

#### GET /api/notifications/device/:deviceID/range
Get notifications within a date range.

Query parameters:
- start: Start date (RFC3339 format)
- end: End date (RFC3339 format)

Example:
```
/api/notifications/device/abc1234/range?start=2024-03-01T00:00:00Z&end=2024-03-31T23:59:59Z
```

#### GET /api/notifications/device/:deviceID/search
Search notifications by title, message, or from field.

Query parameters:
- q: Search query

Example:
```
/api/notifications/device/abc1234/search?q=important
```

#### GET /api/devices
Get a list of all unique device IDs.

## Frontend

The frontend is built using:
- Alpine.js for interactivity
- Tailwind CSS for styling
- Native JavaScript for date handling

Features:
1. Device selection
2. Notification list view
3. Date range filtering
4. Search functionality
5. Responsive design

The web interface is accessible at `http://localhost:8080` 