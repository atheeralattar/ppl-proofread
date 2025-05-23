# Text Proofreader and Content Review

A web application that provides text proofreading and content review using local LLMs (Large Language Models) through Ollama. The application supports multiple models including Llama 3 and Deepseek.

## Features

- **Text Proofreading**: Corrects grammar, spelling, and punctuation errors
- **Content Review**: Provides recommendations for improving content structure and clarity
- **Multiple LLM Models**: Support for different models through Ollama
  - Llama 3
  - Deepseek
- **Model Selection**: Easy switching between different models via dropdown
- **Performance Metrics**: Real-time tracking of request processing times
- **Modern UI**: Clean and responsive interface
- **Markdown Support**: Results are formatted with proper markdown rendering

## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Ollama (will be installed via Docker)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-proof
```

2. Install the required Go dependencies:
```bash
go mod init go-proof
go mod tidy
```

## Running with Docker

1. Start the Ollama service using Docker Compose:
```bash
docker compose up -d
```

2. Pull the required LLM models (this might take a few minutes depending on your internet connection):
```bash
# Pull Llama 3
docker compose exec ollama ollama pull llama3:latest

# Pull Deepseek
docker compose exec ollama ollama pull deepseek:latest
```

3. Start the Go server:
```bash
go run main.go
```

4. Open your web browser and navigate to:
```
http://localhost:8080
```

## Project Structure

- `main.go`: Go server implementation with API endpoints
- `llms/`: Directory containing LLM client implementations
  - `ollama.go`: Ollama client implementation
- `index.html`: Frontend interface with styling and JavaScript
- `prompts/`: Directory containing prompt templates
  - `proofread.txt`: Template for proofreading requests
  - `review.txt`: Template for content review requests
- `docker-compose.yml`: Docker configuration for Ollama service
- `README.md`: Project documentation

## Usage

1. Select your preferred model from the dropdown menu
2. Enter your text in the text area
3. Choose one of the following options:
   - Click "Proofread" to correct grammar and spelling
   - Click "Content Review" to get content improvement recommendations
4. View the formatted results below the buttons

## Performance Monitoring

The application provides real-time performance metrics for each request:
- Generation time: Time taken by the LLM to generate the response
- Total request time: Complete processing time including network overhead
- Response length: Size of the generated response in characters

These metrics are displayed in the server console for monitoring and optimization.

## API Endpoints

- `POST /proofread`: Proofreads the submitted text
  - Request body: `{"text": "your text", "model": "model-name"}`
- `POST /review`: Provides content review and recommendations
  - Request body: `{"text": "your text", "model": "model-name"}`

## Available Models

The application supports the following models through Ollama:
- `llama3:latest`: Llama 3 model for general text processing
- `deepseek:latest`: Deepseek model optimized for code and technical content

## Troubleshooting

1. If models fail to load:
   - Check if Ollama service is running: `docker compose ps`
   - Verify model installation: `docker compose exec ollama ollama list`
   - Pull models again if needed

2. If the server fails to start:
   - Ensure port 8080 is available
   - Check if Ollama is running on port 11434
   - Verify Docker containers are running

3. If requests are timing out:
   - Check your system resources
   - Adjust the timeout in `main.go` if needed (default: 180 seconds)
   - Ensure the selected model is properly loaded

## Contributing

Feel free to submit issues and enhancement requests!