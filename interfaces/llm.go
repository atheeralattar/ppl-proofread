package interfaces

type LLM interface {
	Proofread(doc string) (string, error)
	Review(doc string) (string, error)
}
