package vizio_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grandcat/zeroconf"
	"log"
	"strconv"
	"time"
)

type PairingStartResponse struct {
	STATUS struct {
		RESULT string `json:"RESULT"`
		DETAIL string `json:"DETAIL"`
	} `json:"STATUS"`
	ITEM struct {
		CHALLENGE_TYPE    int `json:"CHALLENGE_TYPE"`
		PAIRING_REQ_TOKEN int `json:"PAIRING_REQ_TOKEN"`
	} `json:"ITEM"`
}
type VizioSetup struct {
	PairingStartResponse
}

type VizioDevice struct {
	Name  string
	IP    string
	Port  int
	Model string
	ID    string
}

func (s *VizioSetup) ListDevices() []VizioDevice {
	services := []VizioDevice{}

	appendService := func(service *zeroconf.ServiceEntry) {
		model := service.Text[0]
		id := service.Text[1]

		// handle id decode for various discovered use cases
		if _, err := strconv.ParseInt(id, 16, 64); err != nil {
			id = fmt.Sprintf("%x", id)
		}

		device := VizioDevice{service.Instance, service.AddrIPv4[0].String(), service.Port, model, id}
		services = append(services, device)
	}

	resolver, err := zeroconf.NewResolver()
	if err != nil {
		fmt.Println("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			appendService(entry)
		}
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err = resolver.Browse(ctx, "_viziocast._tcp.", "local.", entries)
	if err != nil {
		fmt.Println("Failed to browse:", err.Error())
	}

	<-ctx.Done()

	return services
}
func (s *VizioSetup) Pair(device VizioDevice) error {
	url := fmt.Sprintf("https://%s:%d", device.IP, device.Port)

	response, err := sendRequest(url, "PUT", "/pairing/start", "", map[string]string{
		"DEVICE_NAME": device.Name,
		"DEVICE_ID":   device.ID,
	})
	if err != nil {
		return err
	}

	s.PairingStartResponse = PairingStartResponse{}

	err = json.Unmarshal(response, &s.PairingStartResponse)
	if err != nil {
		log.Printf("failed to parse PairingStartResponse json: %e\n", err)
		return err
	}
	return nil
}
func (s *VizioSetup) PostChallenge(device VizioDevice, code int) (string, error) {
	url := fmt.Sprintf("https://%s:%d", device.IP, device.Port)
	type pairingPair struct {
		DEVICE_ID         string `json:"DEVICE_ID"`
		CHALLENGE_TYPE    int    `json:"CHALLENGE_TYPE"`
		RESPONSE_VALUE    string `json:"RESPONSE_VALUE"`
		PAIRING_REQ_TOKEN int    `json:"PAIRING_REQ_TOKEN"`
	}

	response, err := sendRequest(url, "PUT", "/pairing/pair", "", pairingPair{
		device.ID, s.ITEM.CHALLENGE_TYPE, strconv.Itoa(code), s.ITEM.PAIRING_REQ_TOKEN,
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	type SubmitChallengeResponse struct {
		ITEM struct {
			AUTH_TOKEN string
		}
	}
	output := SubmitChallengeResponse{}

	err = json.Unmarshal(response, &output)
	if err != nil {
		log.Printf("failed to parse outout in SubmitChallenge: %e", err)
		return "", err
	}
	return output.ITEM.AUTH_TOKEN, nil
}
