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

// main function orchestrates reading, processing, and writing output.
func main() {
	inputFile := "input.txt"
	outputFile := "output.json"

	prompts, err := readPrompts(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	var outputData []Response

	for _, prompt := range prompts {
		timeSent := time.Now().Unix()  // Current UNIX timestamp for when the prompt was sent
		message := generateResponse(prompt)
		timeRecvd := time.Now().Unix()  // Current UNIX timestamp for when the response was received

		// Create the JSON object
		jsonObject := Response{
			Prompt:   strings.TrimSpace(prompt),
			Message:  message,
			TimeSent: timeSent,
			TimeRecvd: timeRecvd,
			Source:   source,
		}

		// Add to output list
		outputData = append(outputData, jsonObject)
	}

	// Write to output.json file
	outputJSON, err := json.MarshalIndent(outputData, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	err = os.WriteFile(outputFile, outputJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Printf("Output successfully written to %s\n", outputFile)
}
