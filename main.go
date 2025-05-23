package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-proof/llms"
	"go-proof/prompts"
	"log"
	"net/http"
	"time"
)

type RequestBody struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

type ResponseBody struct {
	Text string `json:"text"`
}

var ollamaClient *llms.OllamaClient

func main() {
	// Initialize Ollama client with default model
	// The model will be changed based on user selection
	ollamaClient = llms.NewOllamaClient("http://localhost:11434", "llama3:latest",
		llms.WithTimeout(180*time.Second),
	)

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir(".")))

	// API endpoints
	http.HandleFunc("/proofread", handleProofread)
	http.HandleFunc("/review", handleReview)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProofread(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the model based on user selection
	ollamaClient.SetModel(reqBody.Model)

	// Create the prompt with the user's text
	prompt := fmt.Sprintf(prompts.ProofreadPrompt, reqBody.Text)
	
	// Generate response using Ollama
	ctx := context.Background()
	genStartTime := time.Now()
	response, err := ollamaClient.Generate(ctx, prompt)
	genDuration := time.Since(genStartTime)

	if err != nil {
		log.Printf("Error generating proofread response: %v", err)
		http.Error(w, "Failed to generate content", http.StatusInternalServerError)
		return
	}

	// Send the response
	resp := ResponseBody{Text: response}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	totalDuration := time.Since(startTime)
	fmt.Printf("Generation time: %v\n", genDuration)
	fmt.Printf("Total request time: %v\n", totalDuration)
	fmt.Printf("Response length: %d characters\n\n", len(response))
}

func handleReview(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the model based on user selection
	ollamaClient.SetModel(reqBody.Model)

	// Create the prompt with the user's text
	prompt := fmt.Sprintf(prompts.ReviewPrompt, reqBody.Text)
	fmt.Printf("\n[%s] Processing review request for model: %s\n", time.Now().Format("15:04:05"), reqBody.Model)
	fmt.Printf("Prompt: %s\n", prompt)

	// Generate response using Ollama
	ctx := context.Background()
	genStartTime := time.Now()
	response, err := ollamaClient.Generate(ctx, prompt)
	genDuration := time.Since(genStartTime)

	if err != nil {
		log.Printf("Error generating review response: %v", err)
		http.Error(w, "Failed to generate content", http.StatusInternalServerError)
		return
	}

	// Send the response
	resp := ResponseBody{Text: response}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	totalDuration := time.Since(startTime)
	fmt.Printf("Generation time: %v\n", genDuration)
	fmt.Printf("Total request time: %v\n", totalDuration)
	fmt.Printf("Response length: %d characters\n\n", len(response))
}
