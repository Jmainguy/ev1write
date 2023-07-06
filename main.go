package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

type Card struct {
	Card struct {
		Counter0  string `json:"Counter0"`
		Counter1  string `json:"Counter1"`
		Counter2  string `json:"Counter2"`
		Signature string `json:"Signature"`
		TBO0      string `json:"TBO_0"`
		TBO1      string `json:"TBO_1"`
		Tearing0  string `json:"Tearing0"`
		Tearing1  string `json:"Tearing1"`
		Tearing2  string `json:"Tearing2"`
		UID       string `json:"UID"`
		Version   string `json:"Version"`
	}
	Blocks struct {
		Zero  string `json:"0"`
		One   string `json:"1"`
		One0  string `json:"10"`
		One1  string `json:"11"`
		One2  string `json:"12"`
		One3  string `json:"13"`
		One4  string `json:"14"`
		One5  string `json:"15"`
		One6  string `json:"16"`
		One7  string `json:"17"`
		One8  string `json:"18"`
		One9  string `json:"19"`
		Two   string `json:"2"`
		Three string `json:"3"`
		Four  string `json:"4"`
		Five  string `json:"5"`
		Six   string `json:"6"`
		Seven string `json:"7"`
		Eight string `json:"8"`
		Nine  string `json:"9"`
	} `json:"blocks"`
}

func main() {
	// Check if filename is provided as the first argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a JSON filename as the first argument")
	}

	// Read the JSON file
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err)
	}

	// Parse the JSON data into Card struct
	var card Card
	err = json.Unmarshal(data, &card)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	// Create a slice to store commands
	commands := make([]string, 0)

	// Add signature command
	signature := card.Card.Signature
	sigString := fmt.Sprintf("script run hf_mfu_magicwrite -s %s", signature)
	commands = append(commands, sigString)

	// Iterate over fields of card.Blocks
	blockFields := reflect.ValueOf(card.Blocks)
	for i := 0; i < blockFields.NumField(); i++ {
		key := blockFields.Type().Field(i).Tag.Get("json")
		value := blockFields.Field(i).String()

		if value == "" {
			continue // Skip empty values
		}

		cmdString := fmt.Sprintf("hf mfu wrbl -b %s -d %s", key, value)
		commands = append(commands, cmdString)
	}

	// Execute the commands
	runCommands(commands)
}

func runCommands(commands []string) {
	for _, cmd := range commands {
		for {
			output, err := exec.Command("pm3", "-c", cmd).CombinedOutput()
			if err != nil {
				log.Printf("Command failed: %s", err)
				break
			}
			if strings.Contains(string(output), "A2 Cmd failed. Card timeout.") {
				log.Printf("Retrying after 1 second...")
				time.Sleep(1 * time.Second)
			} else {
				log.Printf("Command output:\n%s\n", output)
				break
			}
		}
	}
}
