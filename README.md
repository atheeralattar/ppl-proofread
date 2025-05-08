# Text Proofreader and Content Review

Built with Go and the Gemini API. This application provides a user-friendly interface for proofreading text and getting content improvement recommendations.

## Features

- **Text Proofreading**: Corrects grammar, spelling, and punctuation errors
- **Content Review**: Provides recommendations for improving content structure and clarity
- **Customizable Prompts**: Easy to modify prompts through text files
- **Modern UI**: Clean and responsive interface
- **Markdown Support**: Results are formatted with proper markdown rendering

## Prerequisites

- Go 1.16 or higher
- Google Gemini API key

## Installation

1. Clone the repository:
```bash
git clone <https://github.com/atheeralattar/ppl-proofread>
cd go-proof
```

2. Install the required Go dependencies:
```bash
go mod init go-proof
go get google.golang.org/genai
```

3. Update the API key (if needed) in `main.go`:
Replace the placeholder API key with your own Gemini API key:
```go
APIKey: "your-api-key-here"
```

## Project Structure

- `main.go`: Go server implementation with API endpoints
- `index.html`: Frontend interface with styling and JavaScript
- `prompts/`: Directory containing prompt templates
  - `proofread.txt`: Template for proofreading requests
  - `review.txt`: Template for content review requests
- `README.md`: Project documentation

## Customizing Prompts

You can customize the prompts used by the application by editing the text files in the `prompts/` directory:

- `prompts/proofread.txt`: Modify the proofreading prompt
- `prompts/review.txt`: Modify the content review prompt

Each prompt file should contain a single `%s` placeholder where the user's text will be inserted.

## Running the Application

1. Start the Go server:
```bash
go run main.go
```

2. Open your web browser and navigate to:
```
http://localhost:8080
```

## Usage

1. Enter your text in the text area
2. Choose one of the following options:
   - Click "Proofread" to correct grammar and spelling
   - Click "Content Review" to get content improvement recommendations
3. View the formatted results below the buttons

## API Endpoints

- `POST /proofread`: Proofreads the submitted text
- `POST /review`: Provides content review and recommendations

