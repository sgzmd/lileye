#!/bin/bash

# Device 1: Personal Phone
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"WhatsApp Message","message":"Hey, how are you?","timestamp":"2024-03-20T10:00:00Z","package_name":"com.whatsapp","from":"Alice","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Gmail","message":"Meeting tomorrow","timestamp":"2024-03-20T10:15:00Z","package_name":"com.google.gmail","from":"boss@company.com","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Calendar Reminder","message":"Dentist appointment","timestamp":"2024-03-20T11:00:00Z","package_name":"com.google.calendar","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Telegram","message":"New group message","timestamp":"2024-03-20T11:30:00Z","package_name":"org.telegram.messenger","from":"Family Group","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Battery Low","message":"20% battery remaining","timestamp":"2024-03-20T12:00:00Z","package_name":"android.system","device_id":"phone1"}'

# Device 2: Work Phone
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Slack","message":"New message in #general","timestamp":"2024-03-20T09:00:00Z","package_name":"com.slack","from":"John","device_id":"phone2"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Microsoft Teams","message":"Team meeting starting","timestamp":"2024-03-20T09:30:00Z","package_name":"com.microsoft.teams","device_id":"phone2"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Outlook","message":"Project deadline reminder","timestamp":"2024-03-20T10:00:00Z","package_name":"com.microsoft.outlook","from":"Project Manager","device_id":"phone2"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Jira","message":"Task assigned to you","timestamp":"2024-03-20T10:30:00Z","package_name":"com.atlassian.jira","device_id":"phone2"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Google Drive","message":"Document shared with you","timestamp":"2024-03-20T11:00:00Z","package_name":"com.google.drive","from":"colleague@company.com","device_id":"phone2"}'

# Device 3: Tablet
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Netflix","message":"New episode available","timestamp":"2024-03-20T18:00:00Z","package_name":"com.netflix.android","device_id":"tablet1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Kindle","message":"Reading goal achieved","timestamp":"2024-03-20T19:00:00Z","package_name":"com.amazon.kindle","device_id":"tablet1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"YouTube","message":"Channel uploaded new video","timestamp":"2024-03-20T20:00:00Z","package_name":"com.google.youtube","from":"TechChannel","device_id":"tablet1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Spotify","message":"New playlist suggestion","timestamp":"2024-03-20T21:00:00Z","package_name":"com.spotify.music","device_id":"tablet1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Weather Alert","message":"Rain expected tomorrow","timestamp":"2024-03-20T22:00:00Z","package_name":"com.weather.app","device_id":"tablet1"}'

# Previous Day Notifications (Phone 1)
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Instagram","message":"New follower","timestamp":"2024-03-19T10:00:00Z","package_name":"com.instagram.android","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Facebook","message":"Birthday reminder","timestamp":"2024-03-19T11:00:00Z","package_name":"com.facebook.katana","from":"Events","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Twitter","message":"Trending in your area","timestamp":"2024-03-19T12:00:00Z","package_name":"com.twitter.android","device_id":"phone1"}'

# Previous Day Notifications (Phone 2)
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Calendar","message":"Team lunch tomorrow","timestamp":"2024-03-19T15:00:00Z","package_name":"com.google.calendar","device_id":"phone2"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Slack","message":"New channel invitation","timestamp":"2024-03-19T16:00:00Z","package_name":"com.slack","from":"Team Lead","device_id":"phone2"}'

# Previous Day Notifications (Tablet)
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Prime Video","message":"Continue watching","timestamp":"2024-03-19T20:00:00Z","package_name":"com.amazon.primevideo","device_id":"tablet1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Game Update","message":"New level unlocked","timestamp":"2024-03-19T21:00:00Z","package_name":"com.game.example","device_id":"tablet1"}'

# System Notifications
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"System Update","message":"Android update available","timestamp":"2024-03-20T09:00:00Z","package_name":"android.system","device_id":"phone1"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Security Alert","message":"New device signed in","timestamp":"2024-03-20T10:00:00Z","package_name":"com.google.android.gms","device_id":"phone2"}'
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{"title":"Storage Alert","message":"Storage space low","timestamp":"2024-03-20T11:00:00Z","package_name":"android.system","device_id":"tablet1"}' 