#!/bin/bash

# Get current date and calculate date range
CURRENT_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
TWO_WEEKS_AGO=$(date -u -v-14d +"%Y-%m-%dT%H:%M:%SZ")
TWO_WEEKS_FROM_NOW=$(date -u -v+14d +"%Y-%m-%dT%H:%M:%SZ")

# Function to generate a random time within a date
generate_random_time() {
    local date=$1
    local hour=$((RANDOM % 24))
    local minute=$((RANDOM % 60))
    local second=$((RANDOM % 60))
    echo "${date}T${hour}:${minute}:${second}Z"
}

# Function to send a notification
send_notification() {
    local title=$1
    local message=$2
    local timestamp=$3
    local package_name=$4
    local from=$5
    local device_id=$6

    curl -X POST http://localhost:8080/api/notifications \
        -H "Content-Type: application/json" \
        -d "{
            \"title\": \"$title\",
            \"message\": \"$message\",
            \"timestamp\": \"$timestamp\",
            \"package_name\": \"$package_name\",
            \"from\": \"$from\",
            \"device_id\": \"$device_id\"
        }"
}

# Generate notifications for different devices and apps
DEVICES=("phone1" "phone2" "tablet1")
MESSAGING_APPS=(
    "com.whatsapp" "WhatsApp"
    "com.facebook.orca" "Messenger"
    "com.telegram" "Telegram"
    "com.snapchat.android" "Snapchat"
    "com.instagram.android" "Instagram"
)

EMAIL_APPS=(
    "com.google.android.gm" "Gmail"
    "com.microsoft.office.outlook" "Outlook"
    "com.yahoo.mobile.client.android.mail" "Yahoo Mail"
)

SYSTEM_APPS=(
    "com.android.systemui" "System"
    "com.google.android.apps.nexuslauncher" "Launcher"
    "com.android.settings" "Settings"
)

ENTERTAINMENT_APPS=(
    "com.netflix.mediaclient" "Netflix"
    "com.spotify.music" "Spotify"
    "com.google.android.youtube" "YouTube"
    "com.amazon.avod.thirdpartyclient" "Prime Video"
)

# Generate notifications for each device
for device in "${DEVICES[@]}"; do
    echo "Generating notifications for $device..."
    
    # Generate notifications for each day in the range
    current=$TWO_WEEKS_AGO
    while [ "$current" \< "$TWO_WEEKS_FROM_NOW" ]; do
        # Generate 2-5 notifications per day
        num_notifications=$((RANDOM % 4 + 2))
        
        for ((i=1; i<=num_notifications; i++)); do
            # Generate random time for this notification
            timestamp=$(generate_random_time "${current%%T*}")
            
            # Randomly select an app category
            app_category=$((RANDOM % 4))
            
            case $app_category in
                0) # Messaging apps
                    app_index=$((RANDOM % ${#MESSAGING_APPS[@]} / 2))
                    package_name=${MESSAGING_APPS[$((app_index * 2))]}
                    app_name=${MESSAGING_APPS[$((app_index * 2 + 1))]}
                    from=("John" "Alice" "Bob" "Charlie" "Diana" "Emma")
                    from_name=${from[$((RANDOM % ${#from[@]}))]}
                    send_notification "New message from $from_name" "Hey, how are you?" "$timestamp" "$package_name" "$from_name" "$device"
                    ;;
                    
                1) # Email apps
                    app_index=$((RANDOM % ${#EMAIL_APPS[@]} / 2))
                    package_name=${EMAIL_APPS[$((app_index * 2))]}
                    app_name=${EMAIL_APPS[$((app_index * 2 + 1))]}
                    from=("boss@company.com" "hr@company.com" "team@project.com" "support@service.com")
                    from_name=${from[$((RANDOM % ${#from[@]}))]}
                    send_notification "New email from $from_name" "Important update about the project" "$timestamp" "$package_name" "$from_name" "$device"
                    ;;
                    
                2) # System apps
                    app_index=$((RANDOM % ${#SYSTEM_APPS[@]} / 2))
                    package_name=${SYSTEM_APPS[$((app_index * 2))]}
                    app_name=${SYSTEM_APPS[$((app_index * 2 + 1))]}
                    send_notification "System Update Available" "A new system update is ready to install" "$timestamp" "$package_name" "System" "$device"
                    ;;
                    
                3) # Entertainment apps
                    app_index=$((RANDOM % ${#ENTERTAINMENT_APPS[@]} / 2))
                    package_name=${ENTERTAINMENT_APPS[$((app_index * 2))]}
                    app_name=${ENTERTAINMENT_APPS[$((app_index * 2 + 1))]}
                    send_notification "New content available" "Check out the latest releases" "$timestamp" "$package_name" "$app_name" "$device"
                    ;;
            esac
            
            # Add a small delay between notifications
            sleep 0.5
        done
        
        # Move to next day
        current=$(date -u -v+1d -j -f "%Y-%m-%dT%H:%M:%SZ" "$current" +"%Y-%m-%dT%H:%M:%SZ")
    done
done

echo "Test data generation complete!" 