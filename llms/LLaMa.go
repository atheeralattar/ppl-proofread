package llms

import (
	"fmt"
	)


type LLaMa struct {
	Client string
}

func (o *LLaMa) Proofread (doc string) (string, error) {

	fmt.Println("Proofrading using LLaMA...")
	
	return doc, nil
	
	}

func (o *LLaMa) Review (doc string) (string, error) {
	fmt.Println("Reviewing using LLaMA...")
	return doc, nil
}
