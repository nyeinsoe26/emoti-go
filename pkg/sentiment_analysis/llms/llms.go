package llms

import "github.com/nyeinsoe26/emoti-go/pkg/models"

const OpenAI = "openai"

type LLM interface {
	GetSentiments(req models.Request) (*models.SentimentAnalysis, error)
}
