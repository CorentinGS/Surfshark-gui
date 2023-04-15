package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:embed scripts/connect.sh
var connectScript string

//go:embed scripts/wireguard.sh
var connectWireguardScript string

var (
	itemsCountries = make(map[string][]string, 0)
	itemsWireguard = make(map[string][]string, 0)
)

type item struct {
	country  string
	city     string
	fileName string
}

func Connect(fileName string) error {
	// Connect to the VPN server.
	cmd := exec.Command("sh", "-c", "FILE="+fileName+";USERNAME="+os.Getenv("VPN_USERNAME")+";PASSWORD="+os.Getenv("VPN_PASSWORD")+";"+connectScript)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Error connecting to VPN server: %v", err)
	}
	log.Printf("Connected")
	return nil
}

func ConnectWireguard(fileName string) error {
	// Connect to the VPN server.
	cmd := exec.Command("sh", "-c", "FILE="+fileName+";"+connectWireguardScript)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Error connecting to Wireguard server: %v", err)
	}
	log.Printf("Connected")
	return nil
}

func ListItems() []item {
	dir := "/etc/openvpn"
	files, _ := os.ReadDir(dir)
	var items []item
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, "_udp.ovpn") {
			continue
		}
		country := fileName[:2]
		city := fileName[3:6]
		items = append(items, item{country: country, city: city, fileName: fileName})
	}

	return items
}

func ListWireguard() []item {
	dir := "/etc/wireguard"
	files, _ := os.ReadDir(dir)
	var items []item
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".conf") {
			continue
		}
		country := fileName[:2]
		city := fileName[3:6]
		items = append(items, item{country: country, city: city, fileName: fileName})
	}

	return items
}

func FillItemsCountry(items []item) {
	// list of countries
	for _, myItem := range items {
		itemsCountries[myItem.country] = append(itemsCountries[myItem.country], myItem.fileName)
	}
}

func FillItemsWireguard(items []item) {
	// list of countries
	for _, myItem := range items {
		itemsWireguard[myItem.country] = append(itemsWireguard[myItem.country], myItem.fileName)
	}
}

func ListItemsPathsByCountry(country string) []string {
	return itemsCountries[country]
}

func ListItemsPathsByCountryWireguard(country string) []string {
	return itemsWireguard[country]
}
