// Package prompts provides functionality for reading and managing prompt templates
// used in the text proofreading and review application.
package prompts

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	// Global variables to store the prompt templates
	ProofreadPrompt string
	ReviewPrompt    string
)

// init initializes the prompt templates by reading them from files
func init() {
	var err error
	ProofreadPrompt, err = ReadPromptFile("proofread.txt")
	if err != nil {
		log.Fatal("Failed to read proofread prompt:", err)
	}

	ReviewPrompt, err = ReadPromptFile("review.txt")
	if err != nil {
		log.Fatal("Failed to read review prompt:", err)
	}
}

// ReadPromptFile reads a prompt template from the specified file.
// It returns the content as a string and any error encountered.
// The filename should be relative to the prompts directory.
func ReadPromptFile(filename string) (string, error) {
	// Ensure we're reading from the prompts directory
	fullPath := filepath.Join("prompts", filename)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("error reading prompt file %s: %w", filename, err)
	}

	return string(content), nil
}

// GetPrompts returns the current proofread and review prompts.
// This is useful if you need to access the prompts after initialization.
func GetPrompts() (string, string) {
	return ProofreadPrompt, ReviewPrompt
}
