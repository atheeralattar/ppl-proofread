# Text Proofreader Web Application

A web-based text proofreading and content review application built with Go and the Gemini API. This application provides a user-friendly interface for proofreading text and getting content improvement recommendations.

## Features

- **Text Proofreading**: Corrects grammar, spelling, and punctuation errors
- **Content Review**: Provides recommendations for improving content structure and clarity
- **Modern UI**: Clean and responsive interface
- **Markdown Support**: Results are formatted with proper markdown rendering
- **Real-time Processing**: Immediate feedback with loading indicators

## Prerequisites

- Go 1.16 or higher
- Google Gemini API key
- Web browser with JavaScript enabled

## Installation

1. Clone the repository:
```bash
git clone <your-repository-url>
cd go-proof
```

2. Install the required Go dependencies:
```bash
go mod init go-proof
go get google.golang.org/genai
```

3. Update the API key in `main.go`:
Replace the placeholder API key with your own Gemini API key:
```go
APIKey: "your-api-key-here"
```

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

## Project Structure

- `main.go`: Go server implementation with API endpoints
- `index.html`: Frontend interface with styling and JavaScript
- `README.md`: Project documentation

## API Endpoints

- `POST /proofread`: Proofreads the submitted text
- `POST /review`: Provides content review and recommendations

## Error Handling

The application includes error handling for:
- Empty text submissions
- API connection issues
- Invalid requests
- Server errors

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is licensed under the MIT License - see the LICENSE file for details. 