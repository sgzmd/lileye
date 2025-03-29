#!/bin/bash

# ASCII art warning
echo -e "\033[1;31m"
cat << "EOF"
    ⚠️  WARNING  ⚠️
    ⚠️  WARNING  ⚠️
    ⚠️  WARNING  ⚠️

███████╗██╗   ██╗ ██████╗██╗  ██╗██╗███╗   ██╗ ██████╗ 
██╔════╝██║   ██║██╔════╝██║  ██║██║████╗  ██║██╔══██╗
███████╗██║   ██║██║     ███████║██║██╔██╗ ██║██║  ██║
╚════██║██║   ██║██║     ██╔══██║██║██║╚██╗██║██║  ██║
███████║╚██████║╚██████╗██║  ██║██║██║ ╚████║██████╔╝
╚══════╝ ╚═════╝ ╚═════╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚═════╝ 

██╗    ██╗██╗███╗   ██╗██╗██╗██╗██╗██╗██╗██╗██╗██╗██╗██╗
██║    ██║██║████╗  ██║██║██║██║██║██║██║██║██║██║██║██║
██║ █╗ ██║██║██╔██╗ ██║██║██║██║██║██║██║██║██║██║██║██║
██║███╗██║██║██║╚██╗██║██║██║██║██║██║██║██║██║██║██║██║
╚███╔███╔╝██║██║ ╚████║██║██║██║██║██║██║██║██║██║██║██║
 ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝╚═╝╚═╝╚═╝╚═╝╚═╝╚═╝╚═╝╚═╝╚═╝╚═╝
EOF
echo -e "\033[0m"

echo -e "\033[1;31m"
echo "THIS SCRIPT WILL DELETE ALL NOTIFICATIONS FROM THE DATABASE!"
echo "THIS ACTION CANNOT BE UNDONE!"
echo -e "\033[0m"

# Ask for confirmation
echo -e "\033[1;33m"
read -p "Are you absolutely sure you want to continue? Type 'YES' to confirm: " confirm
echo -e "\033[0m"

if [ "$confirm" = "YES" ]; then
    echo "Deleting all notifications..."
    curl -X DELETE http://localhost:8080/api/notifications/all
    echo "Done! All notifications have been deleted."
else
    echo "Operation cancelled."
fi 