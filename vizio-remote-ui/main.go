package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	vizio_api "vizio-api-go/vizio-api"
)

func ShowModal(myWindow fyne.Window, device vizio_api.VizioDevice) {
	setup := vizio_api.VizioSetup{}
	log.Println("Pairing device", device.IP)
	setup.Pair(device)
	var pinModal *widget.PopUp
	pinInput := widget.NewEntry()
	//pinInput.SetPlaceHolder("Enter text...")
	closePinModal := widget.NewButton("Close", func() {
		pin, err := strconv.Atoi(pinInput.Text)
		if err != nil {
			log.Printf("Error parsing pin: %e\n", err)
		} else {
			tokenValue, err := setup.PostChallenge(device, pin)
			if err != nil {
				log.Printf("Error getting token: %e\n", err)
			} else {
				err = WriteToken("token", tokenValue, device)
				if err != nil {
					log.Println(err)
				}
			}
		}
		pinModal.Hide()
	})
	pinModal = widget.NewModalPopUp(container.New(layout.NewVBoxLayout(), pinInput, closePinModal), myWindow.Canvas())
	pinModal.Show()
}

var SelectedToken Token
var DPad vizio_api.DPadAPI
var Volume vizio_api.VolumeAPI
var Power vizio_api.PowerAPI

func main() {
	var myWindow fyne.Window

	// Available TVs
	setup := vizio_api.VizioSetup{}
	devices := setup.ListDevices()
	deviceNames := []string{}
	for i := range devices {
		deviceNames = append(deviceNames, devices[i].Name)
	}
	deviceSelector := widget.NewSelect(deviceNames, func(value string) {
		log.Println("Select set to", value)

		if TokenExistForName("token", value) {
			log.Println("Token exists for device")
			SelectedToken = GetTokenByDeviceName("token", value)
			log.Println("Using API", TokenToURL(SelectedToken))

			DPad = vizio_api.DPadAPI{TokenToURL(SelectedToken), SelectedToken.Token}
			Volume = vizio_api.VolumeAPI{TokenToURL(SelectedToken), SelectedToken.Token}
			Power = vizio_api.PowerAPI{TokenToURL(SelectedToken), SelectedToken.Token}

			log.Println("Registered device")
		} else {
			log.Println("Creating new token for device")
			ShowModal(myWindow, GetDeviceByName(value, devices))
		}

	})

	upBtn := widget.NewButton("UP", func() {
		DPad.Up()
	})
	downBtn := widget.NewButton("DOWN", func() {
		DPad.Down()
	})
	leftBtn := widget.NewButton("LEFT", func() {
		DPad.Left()
	})
	rightBtn := widget.NewButton("RIGHT", func() {
		DPad.Right()
	})
	selectBtn := widget.NewButton("SELECT", func() {
		DPad.Select()
	})
	backBtn := widget.NewButton("BACK", func() {
		DPad.Back()
	})
	volUpBtn := widget.NewButton("+", func() {
		Volume.VolumeUp()
	})
	volDownBtn := widget.NewButton("-", func() {
		Volume.VolumeDown()
	})
	powerBtn := widget.NewButton("POWER", func() {
		Power.PowerToggle()
	})

	myApp := app.New()
	myWindow = myApp.NewWindow("Box Layout")

	grid := container.New(layout.NewGridLayout(3),
		powerBtn, layout.NewSpacer(), layout.NewSpacer(),
		layout.NewSpacer(), upBtn, layout.NewSpacer(),
		leftBtn, selectBtn, rightBtn,
		layout.NewSpacer(), downBtn, layout.NewSpacer(),
		layout.NewSpacer(), backBtn, layout.NewSpacer(),
		layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(),
		layout.NewSpacer(), layout.NewSpacer(), volUpBtn,
		layout.NewSpacer(), layout.NewSpacer(), volDownBtn,
	)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), deviceSelector, grid))
	myWindow.ShowAndRun()

}
