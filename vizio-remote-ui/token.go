package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"vizio-api-go/vizio-api"
)

type Token struct {
	IP    string `json:"IP"`
	Port  int    `json:"Port"`
	ID    string `json:"ID"`
	Token string `json:"Token"`
	Name  string `json:"Name"`
}

func TokenToURL(token Token) string {
	return fmt.Sprintf("https://%s:%d", token.IP, token.Port)
}

func ReadToken(name string) ([]Token, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return []Token{}, err
	}
	var tokens []Token
	err = json.Unmarshal(data, &tokens)
	if err != nil {
		return []Token{}, err
	}
	return tokens, nil
}

func WriteToken(name string, token string, device vizio_api.VizioDevice) error {
	tokens, _ := ReadToken(name)

	data := Token{device.IP, device.Port, device.ID, token, device.Name}
	tokens = append(tokens, data)

	log.Println("WriteToken", tokens)

	jsonData, err := json.Marshal(tokens)
	if err != nil {
		log.Printf("Failed to marshal tokens: %e\n", err)
		return err
	}
	return os.WriteFile(name, jsonData, 0600)
}

func TokenExistForName(name string, deviceName string) bool {
	log.Println("TokenExistForName", name, deviceName)
	tokens, err := ReadToken(name)
	if err != nil {
		log.Printf("Failed to read tokens: %e\n", err)
	}
	log.Println("tokens", tokens)
	for _, token := range tokens {
		log.Println("name in file:", token.Name, "detected device:", deviceName, "same:", token.Name == deviceName)
		if token.Name == deviceName {
			return true
		}
	}
	return false
}

func GetTokenByDeviceName(name string, deviceName string) Token {
	tokens, _ := ReadToken(name)
	for _, token := range tokens {
		if token.Name == deviceName {
			return token
		}
	}
	return Token{}
}
