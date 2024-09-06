package models

// SentimentAnalysis struct to hold sentiment results
type SentimentAnalysis struct {
	Sentiments []string `json:"sentiments"`
}

type Request struct {
	Texts               []string
	Model               string
	Temperature         float32
	SentimentCategories []string
}
