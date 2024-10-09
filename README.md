# DASH_LAB_Assignment

This repository contains a Go-based application that reads prompts from an input file, generates responses, and outputs the results in JSON format. The project supports two modes of operation:

- **Single Output Mode**: Processes all prompts and writes the results into a single `output.json` file.
- **Multi-Client Mode**: Distributes the prompts among multiple clients and generates separate JSON files for each client.

## Features

- Reads prompts from a text file.
- Generates dummy responses for each prompt.
- Outputs results in JSON format.
- Supports multi-client prompt distribution.

## Requirements

- **Go 1.16 or later**: Ensure Go is installed on your system. You can follow the instructions [here](https://golang.org/doc/install) to install Go.
- **An input file (`input.txt`)**: The file should contain one prompt per line.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/hegdemanu/DASH_LAB_Assignment.git
    cd DASH_LAB_Assignment
    ```

2. Ensure Go is installed. If Go is not installed, you can download and install it from [golang.org](https://golang.org).

3. Place your `input.txt` file in the project directory. Each line in the file should contain one prompt.

## Usage

### 1. Single Output Mode
This mode processes all prompts from the input file and writes the results into a single JSON file (`output.json`).

**Command**:
```bash
go run generate_output.go
```
### 2. Multi-Client Mode

This mode distributes the prompts among a specified number of clients and writes a separate JSON file for each client.

**Command:**

```bash
go run generate_output_level2.go <number-of-clients>
```
Replace <number-of-clients> with the number of clients you want to simulate. For example:
```bash
go run generate_output_level2.go 3
```

## Input File Format

The input file (`input.txt`) should contain one prompt per line. For example:
What is the capital of France?
Tell me a joke.
Explain Newton’s first law of motion.

## Output

- **Single Output Mode**: A single JSON file (`output.json`) will be generated, containing all the prompts and their generated responses.
  
- **Multi-Client Mode**: A JSON file will be generated for each client (`output_clientX.json`, where `X` is the client number).

### Example Output

An example of an output JSON file:

```json
[
    {
        "Prompt": "What is the capital of France?",
        "Message": "Response for: What is the capital of France?",
        "TimeSent": 1696870924,
        "TimeRecvd": 1696870925,
        "Source": "ChatGPT",
        "ClientID": 1
    },
    {
        "Prompt": "Tell me a joke.",
        "Message": "Response for: Tell me a joke.",
        "TimeSent": 1696870926,
        "TimeRecvd": 1696870927,
        "Source": "ChatGPT",
        "ClientID": 1
    }
]

```
Error Handling

	•	If the input file cannot be read, an error message will be displayed in the terminal.
	•	If JSON marshaling or file writing fails, appropriate error messages will be shown.

Future Improvements

	•	Implement actual AI-based response generation (currently, responses are dummy placeholders).
	•	Add dynamic client assignment and load balancing.
	•	Enhance logging and error reporting.

License

This project is licensed under the MIT License. See the LICENSE file for more details.
