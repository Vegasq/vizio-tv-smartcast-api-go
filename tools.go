package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	val, err := reader.ReadString('\n')
	val = strings.ReplaceAll(val, "\n", "")
	if err != nil {
		log.Println("Failed to read a line from the reader")
	}
	return val
}
