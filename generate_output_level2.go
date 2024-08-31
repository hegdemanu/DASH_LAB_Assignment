package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Response represents the structure of the JSON object to be output.
type Response struct {
	Prompt   string `json:"Prompt"`
	Message  string `json:"Message"`
	TimeSent int64  `json:"TimeSent"`
	TimeRecvd int64 `json:"TimeRecvd"`
	Source   string `json:"Source"`
	ClientID int    `json:"ClientID"`
}

const source = "ChatGPT"

// readPrompts reads the prompts from a specified input text file.
func readPrompts(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var prompts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prompts = append(prompts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return prompts, nil
}

// generateResponse simulates generating a response for a given prompt.
func generateResponse(prompt string) string {
	// Dummy response generation
	return fmt.Sprintf("Response for: %s", prompt)
}

// assignPromptsToClients distributes prompts among clients.
func assignPromptsToClients(prompts []string, numClients int) map[int][]string {
	clientPrompts := make(map[int][]string)
	for i, prompt := range prompts {
		clientID := i % numClients // Distribute prompts across clients
		clientPrompts[clientID] = append(clientPrompts[clientID], prompt)
	}
	return clientPrompts
}

// main function orchestrates reading, processing, and writing output.
func main() {
	inputFile := "input.txt"
	numClients := 3 // Number of clients

	// Read prompts from file
	prompts, err := readPrompts(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Assign prompts to clients
	clientPrompts := assignPromptsToClients(prompts, numClients)

	// Process each client
	for clientID, prompts := range clientPrompts {
		var outputData []Response

		for _, prompt := range prompts {
			timeSent := time.Now().Unix()  // Current UNIX timestamp for when the prompt was sent
			message := generateResponse(prompt)
			timeRecvd := time.Now().Unix()  // Current UNIX timestamp for when the response was received

			// Determine the source
			source := source // Default to "ChatGPT"
			if clientID != 0 {
				source = "user" // Simulating that only client 0 is the source of truth
			}

			// Create the JSON object
			jsonObject := Response{
				Prompt:   strings.TrimSpace(prompt),
				Message:  message,
				TimeSent: timeSent,
				TimeRecvd: timeRecvd,
				Source:   source,
				ClientID: clientID + 1, // Client ID starts from 1
			}

			// Add to output list
			outputData = append(outputData, jsonObject)
		}

		// Write each client's output to its own JSON file
		outputFile := fmt.Sprintf("output_client%d.json", clientID+1)
		outputJSON, err := json.MarshalIndent(outputData, "", "    ")
		if err != nil {
			fmt.Printf("Error marshalling JSON for client %d: %v\n", clientID+1, err)
			continue
		}

		err = os.WriteFile(outputFile, outputJSON, 0644)
		if err != nil {
			fmt.Printf("Error writing output file for client %d: %v\n", clientID+1, err)
			continue
		}

		fmt.Printf("Output successfully written to %s\n", outputFile)
	}
}
