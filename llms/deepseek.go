package llms

import (
	"fmt"
)

type DeepSeek struct {
	Client string
}

func (d *DeepSeek) Proofread(doc string) (string, error){


	fmt.Println("Proofreading using DeepSeek...")
	
	return doc, nil}

func (d *DeepSeek) Review(doc string) (string, error){
	fmt.Println("Reviewing using DeepSeek...")
	return doc, nil
}
