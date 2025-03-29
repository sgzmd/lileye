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
├── scripts/             # Utility scripts
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

The server will start on `http://localhost:8080`. You can access the web interface by opening this URL in your browser.

## Testing the Application

### Running Test Data

To populate the database with test notifications, run the provided script:

```bash
go run scripts/load_test_data.go
```

The script supports several command-line flags:
- `-days-before`: Number of days before today to generate notifications (default: 14)
- `-days-after`: Number of days after today to generate notifications (default: 14)
- `-min-per-day`: Minimum number of notifications per day (default: 2)
- `-max-per-day`: Maximum number of notifications per day (default: 5)
- `-delay`: Delay between notifications in milliseconds (default: 500)
- `-server`: Server URL (default: http://localhost:8080)
- `-devices`: Comma-separated list of devices (default: phone1,phone2,tablet1)

Example with custom parameters:
```bash
go run scripts/load_test_data.go -days-before 7 -days-after 0 -min-per-day 5 -max-per-day 10 -devices "pixel7,ipad"
```

The script will generate various types of notifications:
- Social media notifications (WhatsApp, Messenger, Telegram, etc.)
- Email notifications (Gmail, Outlook, Yahoo Mail)
- System notifications
- Entertainment app notifications (Netflix, Spotify, YouTube)

After running the script, you can:
1. Visit `http://localhost:8080` in your browser
2. Use the device selector to switch between devices
3. Use the date picker to view notifications from different days
4. Use the search box to find specific notifications

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