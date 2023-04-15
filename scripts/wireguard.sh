#!/usr/bin/env sh

# cd to the config directory
cd /etc/wireguard/ || exit 1

echo "Running wireguard"

# Add the connection
nmcli connection import type wireguard file "$FILE" || exit 1


exit 0