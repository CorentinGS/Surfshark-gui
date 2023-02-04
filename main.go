package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
	"log"
	"sort"
	"time"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("OpenDNS selection")

	FillItemsCountry(ListItems())

	countries := make([]string, 0, len(itemsCountries))
	for country := range itemsCountries {
		countries = append(countries, country)
	}

	// Sort the countries by name
	sort.Strings(countries) // sort.Strings is in-place and returns nothing

	selectButton := widget.Select{}
	selectButton.PlaceHolder = "Select"
	selectButton.Options = itemsCountries["fr"]
	selectButton.Disable()

	selectCountryButton := widget.NewSelect(countries, func(country string) {
		selectButton.Options = ListItemsPathsByCountry(country)
		selectButton.ClearSelected()
		selectButton.Enable()
		selectButton.Refresh()
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
		err := Connect(selectButton.Selected)
		if err != nil {
			progress.Hide()
			return
		}

		progress.SetValue(1)
		progress.Hide()
	})

	centered := container.New(layout.NewHBoxLayout(), selectCountryButton, &selectButton, layout.NewSpacer(), okButton, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), centered, progress, layout.NewSpacer()))
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
