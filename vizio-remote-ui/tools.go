package main

import vizio_api "vizio-api-go/vizio-api"

func GetDeviceByName(name string, devices []vizio_api.VizioDevice) vizio_api.VizioDevice {
	for _, device := range devices {
		if device.Name == name {
			return device
		}
	}
	return vizio_api.VizioDevice{}
}
