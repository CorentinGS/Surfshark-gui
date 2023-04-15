#!/usr/bin/env sh

# cd to the config directory
cd /etc/openvpn || exit 1

NAME=$(echo "$FILE" | cut -d'.' -f1)

echo "Running Open Vpn"

# Add the connection
nmcli connection import type openvpn file "$FILE" || exit 1

# Get the connection name using the config file name without the extension
CONNECTION=$(nmcli connection show | grep "$NAME" | awk '{print $1}')

# Set the username and password
echo "Setting username and password"
nmcli connection modify "$CONNECTION" +vpn.data connection-type=password || exit 1
nmcli connection modify "$CONNECTION" +vpn.data username="$USERNAME" || exit 1
nmcli connection modify "$CONNECTION" +vpn.secrets password="$PASSWORD" || exit 1

exit 0