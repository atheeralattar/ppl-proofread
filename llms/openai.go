package llms

import (
	"fmt"
)

type OpenAI struct {
	Client string
}

func (o *OpenAI) Proofread(doc string) (string, error) {
	fmt.Println("Proofreading document using OpenAI")
	return doc, nil
}

func (o *OpenAI) Review(doc string) (string, error) {
	fmt.Println("Reviewing document using OpenAI")
	return doc, nil
}
