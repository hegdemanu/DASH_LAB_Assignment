# DASH_LAB_Assignment

## Overview
This Go-based application processes prompts from an input file, simulating responses, and outputs the results in JSON format. It supports both single and multi-client output modes.

## Features
- **Single Output Mode:** Processes all prompts and outputs to `output.json`.
- **Multi-Client Mode:** Distributes prompts across multiple clients, outputting to separate JSON files per client.

## Requirements
- Go 1.16+
- `input.txt` containing prompts.

## Setup
```bash
git clone https://github.com/hegdemanu/DASH_LAB_Assignment.git
cd DASH_LAB_Assignment
# Ensure Go is installed: https://golang.org/doc/install
