package sentiment_analysis

import (
	"strings"

	"github.com/nyeinsoe26/emoti-go/pkg/models"
	"github.com/nyeinsoe26/emoti-go/pkg/sentiment_analysis/llms"
	"github.com/nyeinsoe26/emoti-go/pkg/sentiment_analysis/llms/openai"
)

type SentimentAnalyzer struct {
	llm llms.LLM
}

func New(llm_type, api_key string) *SentimentAnalyzer {
	s := SentimentAnalyzer{}
	var llm llms.LLM
	switch strings.ToLower(llm_type) {
	case llms.OpenAI:
		llm = openai.New(api_key)
	}
	s.llm = llm
	return &s
}

func (s *SentimentAnalyzer) GetSentiments(req models.Request) (*models.SentimentAnalysis, error) {
	return s.llm.GetSentiments(req)
}
