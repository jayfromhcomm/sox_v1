# Voice-Activated ChatGPT CLI Application

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Introduction

This project aims to provide a voice-activated interface for interacting with ChatGPT, enhanced with a visually appealing CLI experience using charm.sh libraries. The application is developed using Python for the backend logic and Go for the CLI interface.

## Features

- Voice recognition for converting user speech to text
- Text-to-speech for converting GPT-3 responses to voice
- GPT-3 API integration for chat functionality
- CLI interface designed using charm.sh libraries
- Optional feature for custom voice commands

## Prerequisites

- Python 3.x
- Go 1.x
- Microphone and Speakers

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/voice-activated-chatgpt-cli.git

# Navigate to the project directory
cd voice-activated-chatgpt-cli

# Install Python dependencies
pip install -r requirements.txt

# Install Go dependencies
go get -v -t -d ./...
```

## Usage

# Run the Python backend
python main.py

# Run the Go-based CLI interface
go run main.go

## Development 
To contribute to this project:

Fork the repository
Create your feature branch (git checkout -b feature/fooBar)
Commit your changes (git commit -am 'Add some fooBar')
Push to the branch (git push origin feature/fooBar)
Create a new Pull Request

## Testing

```bash
# Run Python tests
pytest tests/

# Run Go tests
go test -v ./...
```

## Contributing

## License

