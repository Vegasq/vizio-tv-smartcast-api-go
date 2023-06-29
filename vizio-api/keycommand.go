package vizio_api

import "log"

func KeyCommand(url, token, endpoint string, codeset int, code int, action string) {
	type KeysBody struct {
		KEYLIST []struct {
			CODESET int    `json:"CODESET"`
			CODE    int    `json:"CODE"`
			ACTION  string `json:"ACTION"`
		} `json:"KEYLIST"`
	}
	log.Println("KeyCommand", url, token, endpoint, codeset, code, action)
	_, err := sendRequest(url, "PUT", endpoint, token, KeysBody{
		KEYLIST: []struct {
			CODESET int    `json:"CODESET"`
			CODE    int    `json:"CODE"`
			ACTION  string `json:"ACTION"`
		}{
			{CODESET: codeset, CODE: code, ACTION: action},
		},
	})
	if err != nil {
		log.Println(err)
	}
}
