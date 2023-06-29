package main

import (
	"fmt"
	"log"
	"strconv"
	"vizio-api-go/vizio-api"
)

func GetTokenCLI() (string, vizio_api.VizioDevice) {
	setup := vizio_api.VizioSetup{}

	// Select device -----------------------------------------------------------------

	devices := setup.ListDevices()
	fmt.Println("Select device to pair:")
	for i := range devices {
		fmt.Printf("%d\t%s\n", i, devices[i].Name)
	}
	deviceS := readUserInput()
	deviceI, err := strconv.Atoi(deviceS)
	if err != nil {
		log.Fatalf("Failed to convert %s to int\n", deviceS)
	}
	fmt.Printf("Selected device is: %s", devices[deviceI].Name)

	// Initiate pairing with device -----------------------------------------------------------------

	err = setup.Pair(devices[deviceI])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Enter the code displayed on the TV")
	codeS := readUserInput()
	codeI, err := strconv.Atoi(codeS)
	if err != nil {
		log.Fatalln("Failed to convert %s to int\n", codeS)
	}

	// Submit Challenge -----------------------------------------------------------------

	token, err := setup.PostChallenge(devices[deviceI], codeI)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Auth toke is: %s\n", token)
	return token, devices[deviceI]
}
