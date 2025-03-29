package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Notification struct {
	Title       string    `json:"title"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	PackageName string    `json:"package_name"`
	From        string    `json:"from"`
	DeviceID    string    `json:"device_id"`
	DeviceName  string    `json:"device_name,omitempty"`
}

type AppInfo struct {
	PackageName string
	Name        string
}

var (
	// Command line flags
	daysBefore      = flag.Int("days-before", 14, "Number of days before today to generate notifications")
	daysAfter       = flag.Int("days-after", 14, "Number of days after today to generate notifications")
	minPerDay       = flag.Int("min-per-day", 2, "Minimum number of notifications per day")
	maxPerDay       = flag.Int("max-per-day", 5, "Maximum number of notifications per day")
	delayMs         = flag.Int("delay", 500, "Delay between notifications in milliseconds")
	serverURL       = flag.String("server", "http://localhost:8080", "Server URL")
	selectedDevices = flag.String("devices", "phone1,phone2,tablet1", "Comma-separated list of devices to generate notifications for")

	// App categories
	messagingApps = []AppInfo{
		{"com.whatsapp", "WhatsApp"},
		{"com.facebook.orca", "Messenger"},
		{"com.telegram", "Telegram"},
		{"com.snapchat.android", "Snapchat"},
		{"com.instagram.android", "Instagram"},
	}

	emailApps = []AppInfo{
		{"com.google.android.gm", "Gmail"},
		{"com.microsoft.office.outlook", "Outlook"},
		{"com.yahoo.mobile.client.android.mail", "Yahoo Mail"},
	}

	systemApps = []AppInfo{
		{"com.android.systemui", "System"},
		{"com.google.android.apps.nexuslauncher", "Launcher"},
		{"com.android.settings", "Settings"},
	}

	entertainmentApps = []AppInfo{
		{"com.netflix.mediaclient", "Netflix"},
		{"com.spotify.music", "Spotify"},
		{"com.google.android.youtube", "YouTube"},
		{"com.amazon.avod.thirdpartyclient", "Prime Video"},
	}

	// Sender names
	messagingSenders = []string{"John", "Alice", "Bob", "Charlie", "Diana", "Emma"}
	emailSenders    = []string{"boss@company.com", "hr@company.com", "team@project.com", "support@service.com"}

	// Random source
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func validateFlags() error {
	if *daysBefore < 0 {
		return fmt.Errorf("days-before must be non-negative")
	}
	if *daysAfter < 0 {
		return fmt.Errorf("days-after must be non-negative")
	}
	if *minPerDay < 1 {
		return fmt.Errorf("min-per-day must be at least 1")
	}
	if *maxPerDay < *minPerDay {
		return fmt.Errorf("max-per-day must be greater than or equal to min-per-day")
	}
	if *delayMs < 0 {
		return fmt.Errorf("delay must be non-negative")
	}
	if *serverURL == "" {
		return fmt.Errorf("server URL cannot be empty")
	}
	if *selectedDevices == "" {
		return fmt.Errorf("devices cannot be empty")
	}
	return nil
}

func randomTime(date time.Time) time.Time {
	hour := rnd.Intn(24)
	minute := rnd.Intn(60)
	second := rnd.Intn(60)
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, second, 0, time.UTC)
}

func randomApp(apps []AppInfo) AppInfo {
	return apps[rnd.Intn(len(apps))]
}

func randomSender(senders []string) string {
	return senders[rnd.Intn(len(senders))]
}

func waitForServer(maxRetries int, retryDelay time.Duration) error {
	url := fmt.Sprintf("%s/", *serverURL)
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return nil
			}
		}
		fmt.Printf("Waiting for server to be ready (attempt %d/%d)...\n", i+1, maxRetries)
		time.Sleep(retryDelay)
	}
	return fmt.Errorf("server not ready after %d attempts", maxRetries)
}

func sendNotification(n Notification) error {
	data, err := json.Marshal(n)
	if err != nil {
		return fmt.Errorf("error marshaling notification: %v", err)
	}

	url := fmt.Sprintf("%s/api/notifications", *serverURL)
	
	// Try up to 3 times with exponential backoff
	for i := 0; i < 3; i++ {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			if i < 2 {
				delay := time.Duration(1<<uint(i)) * time.Second
				fmt.Printf("Error sending notification, retrying in %v... (%v)\n", delay, err)
				time.Sleep(delay)
				continue
			}
			return fmt.Errorf("error sending notification: %v", err)
		}
		defer resp.Body.Close()

		// Read and discard response body
		if _, err := io.Copy(io.Discard, resp.Body); err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			if i < 2 {
				delay := time.Duration(1<<uint(i)) * time.Second
				fmt.Printf("Unexpected status code %d, retrying in %v...\n", resp.StatusCode, delay)
				time.Sleep(delay)
				continue
			}
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		return nil
	}

	return fmt.Errorf("failed to send notification after 3 attempts")
}

func generateNotifications() error {
	now := time.Now().UTC()
	startDate := now.AddDate(0, 0, -*daysBefore)
	endDate := now.AddDate(0, 0, *daysAfter)

	// Parse selected devices
	devices := strings.Split(*selectedDevices, ",")
	if len(devices) == 0 {
		return fmt.Errorf("no devices specified")
	}

	for _, device := range devices {
		fmt.Printf("Generating notifications for %s...\n", device)
		
		for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
			// Generate random number of notifications for this day
			numNotifications := rnd.Intn(*maxPerDay-*minPerDay+1) + *minPerDay
			
			for i := 0; i < numNotifications; i++ {
				appCategory := rnd.Intn(4)
				var notification Notification
				
				switch appCategory {
				case 0: // Messaging apps
					app := randomApp(messagingApps)
					sender := randomSender(messagingSenders)
					notification = Notification{
						Title:       fmt.Sprintf("New message from %s", sender),
						Message:     "Hey, how are you?",
						Timestamp:   randomTime(date),
						PackageName: app.PackageName,
						From:        sender,
						DeviceID:    device,
					}
					
				case 1: // Email apps
					app := randomApp(emailApps)
					sender := randomSender(emailSenders)
					notification = Notification{
						Title:       fmt.Sprintf("New email from %s", sender),
						Message:     "Important update about the project",
						Timestamp:   randomTime(date),
						PackageName: app.PackageName,
						From:        sender,
						DeviceID:    device,
					}
					
				case 2: // System apps
					app := randomApp(systemApps)
					notification = Notification{
						Title:       "System Update Available",
						Message:     "A new system update is ready to install",
						Timestamp:   randomTime(date),
						PackageName: app.PackageName,
						From:        "System",
						DeviceID:    device,
					}
					
				case 3: // Entertainment apps
					app := randomApp(entertainmentApps)
					notification = Notification{
						Title:       "New content available",
						Message:     "Check out the latest releases",
						Timestamp:   randomTime(date),
						PackageName: app.PackageName,
						From:        app.Name,
						DeviceID:    device,
					}
				}

				if err := sendNotification(notification); err != nil {
					return fmt.Errorf("error sending notification: %v", err)
				}

				// Add a small delay between notifications
				time.Sleep(time.Duration(*delayMs) * time.Millisecond)
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()

	if err := validateFlags(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Waiting for server at %s to be ready...\n", *serverURL)
	if err := waitForServer(10, time.Second); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Make sure the server is running with: go run cmd/server/main.go")
		os.Exit(1)
	}

	if err := generateNotifications(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Test data generation complete!")
}
