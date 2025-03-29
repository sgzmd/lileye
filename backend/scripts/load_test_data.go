package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Notification struct {
	Title       string    `json:"title"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	PackageName string    `json:"package_name"`
	From        string    `json:"from"`
	DeviceID    string    `json:"device_id"`
}

type AppInfo struct {
	PackageName string
	Name        string
}

var (
	devices = []string{"phone1", "phone2", "tablet1"}
	
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

	messagingSenders = []string{"John", "Alice", "Bob", "Charlie", "Diana", "Emma"}
	emailSenders    = []string{"boss@company.com", "hr@company.com", "team@project.com", "support@service.com"}
)

func randomTime(date time.Time) time.Time {
	hour := rand.Intn(24)
	minute := rand.Intn(60)
	second := rand.Intn(60)
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, second, 0, time.UTC)
}

func randomApp(apps []AppInfo) AppInfo {
	return apps[rand.Intn(len(apps))]
}

func randomSender(senders []string) string {
	return senders[rand.Intn(len(senders))]
}

func sendNotification(n Notification) error {
	data, err := json.Marshal(n)
	if err != nil {
		return fmt.Errorf("error marshaling notification: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/api/notifications", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error sending notification: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func generateNotifications() error {
	now := time.Now().UTC()
	twoWeeksAgo := now.AddDate(0, 0, -14)
	twoWeeksFromNow := now.AddDate(0, 0, 14)

	for _, device := range devices {
		fmt.Printf("Generating notifications for %s...\n", device)
		
		for date := twoWeeksAgo; date.Before(twoWeeksFromNow); date = date.AddDate(0, 0, 1) {
			// Generate 2-5 notifications per day
			numNotifications := rand.Intn(4) + 2
			
			for i := 0; i < numNotifications; i++ {
				appCategory := rand.Intn(4)
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
				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	if err := generateNotifications(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("Test data generation complete!")
} 