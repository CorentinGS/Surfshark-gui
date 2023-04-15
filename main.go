package main

import (
	"log"
	"sort"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("OpenDNS selection")

	FillItemsCountry(ListItems())

	FillItemsWireguard(ListWireguard())

	countries := make([]string, 0, len(itemsCountries))
	for country := range itemsCountries {
		countries = append(countries, country)
	}

	// Sort the countries by name
	sort.Strings(countries) // sort.Strings is in-place and returns nothing

	selectButton := widget.Select{}
	selectButton.PlaceHolder = "Select"
	selectButton.Disable()

	countriesWireguard := make([]string, 0, len(itemsWireguard))
	for country := range itemsWireguard {
		countriesWireguard = append(countriesWireguard, country)
	}

	// Sort the countries by name
	sort.Strings(countriesWireguard) // sort.Strings is in-place and returns nothing

	selectButtonWireguard := widget.Select{}
	selectButtonWireguard.PlaceHolder = "Select"
	selectButtonWireguard.Disable()

	selectCountryButton := widget.NewSelect(countries, func(country string) {
		selectButton.Options = ListItemsPathsByCountry(country)
		selectButton.ClearSelected()
		selectButton.Enable()
		selectButton.Refresh()
	})

	selectCountryButtonWireguard := widget.NewSelect(countriesWireguard, func(country string) {
		selectButtonWireguard.Options = ListItemsPathsByCountryWireguard(country)
		selectButtonWireguard.ClearSelected()
		selectButtonWireguard.Enable()
		selectButtonWireguard.Refresh()
	})

	progress := widget.NewProgressBar()
	progress.Hide()

	okButton := widget.NewButton("OK", func() {
		progress.Show()
		go func() {
			for i := 0.0; i <= 0.8; i += 0.1 {
				time.Sleep(time.Millisecond * 250)
				progress.SetValue(i)
			}
		}()
		go func() {
			err := Connect(selectButton.Selected)
			if err != nil {
				progress.Hide()
				return
			}

			progress.SetValue(1)
			progress.Hide()
		}()
	})

	okButtonWireguard := widget.NewButton("OK", func() {
		progress.Show()
		go func() {
			for i := 0.0; i <= 0.8; i += 0.1 {
				time.Sleep(time.Millisecond * 250)
				progress.SetValue(i)
			}
		}()
		go func() {
			err := ConnectWireguard(selectButtonWireguard.Selected)
			if err != nil {
				progress.Hide()
				return
			}
			progress.SetValue(1)
			progress.Hide()
		}()

	})

	centered := container.New(layout.NewHBoxLayout(), selectCountryButton, &selectButton, layout.NewSpacer(), okButton, layout.NewSpacer())
	centeredWireguard := container.New(layout.NewHBoxLayout(), selectCountryButtonWireguard, &selectButtonWireguard, layout.NewSpacer(), okButtonWireguard, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), centered, centeredWireguard, progress, layout.NewSpacer()))
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
