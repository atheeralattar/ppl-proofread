package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/genai"
)

type RequestBody struct {
	Text string `json:"text"`
}

type ResponseBody struct {
	Text string `json:"text"`
}

var (
	proofreadPrompt string
	reviewPrompt    string
)

func init() {
	// Read prompt files
	proofreadPrompt = readPromptFile("prompts/proofread.txt")
	reviewPrompt = readPromptFile("prompts/review.txt")
}

func readPromptFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading prompt file %s: %v", filename, err)
	}
	return string(content)
}

func main() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir(".")))

	// API endpoints
	http.HandleFunc("/proofread", handleProofread)
	http.HandleFunc("/review", handleReview)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProofread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  "AIzaSyCGQlkyOKgzHZCBqiXdnxZY9LCE9kEcWSo",
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		http.Error(w, "Failed to create client", http.StatusInternalServerError)
		return
	}

	prompt := fmt.Sprintf(proofreadPrompt, reqBody.Text)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to generate content", http.StatusInternalServerError)
		return
	}

	text, err := result.Text()
	if err != nil {
		http.Error(w, "Failed to get text from result", http.StatusInternalServerError)
		return
	}

	response := ResponseBody{Text: text}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  "AIzaSyCGQlkyOKgzHZCBqiXdnxZY9LCE9kEcWSo",
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		http.Error(w, "Failed to create client", http.StatusInternalServerError)
		return
	}

	prompt := fmt.Sprintf(reviewPrompt, reqBody.Text)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to generate content", http.StatusInternalServerError)
		return
	}

	text, err := result.Text()
	if err != nil {
		http.Error(w, "Failed to get text from result", http.StatusInternalServerError)
		return
	}

	response := ResponseBody{Text: text}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
