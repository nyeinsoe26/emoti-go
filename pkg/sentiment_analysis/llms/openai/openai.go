package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nyeinsoe26/emoti-go/pkg/models"
	openai "github.com/sashabaranov/go-openai"
)

const DefaultModel = "gpt-4o"

type OpenAI struct {
	client *openai.Client
}

func New(openai_api_key string) *OpenAI {
	client := openai.NewClient(openai_api_key)
	return &OpenAI{
		client: client,
	}
}

func (c *OpenAI) GetSentiments(req models.Request) (*models.SentimentAnalysis, error) {
	if len(req.SentimentCategories) == 0 {
		req.SentimentCategories = DefaultSentiments
	}
	labels := strings.Join(func() []string {
		s := make([]string, len(req.SentimentCategories))
		for i, v := range req.SentimentCategories {
			s[i] = fmt.Sprintf("- %s", v)
		}
		return s
	}(), "\n")

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       req.Model,
			Temperature: req.Temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(PromptTemplate, ResponseFormat, labels, req.Texts),
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	var res models.SentimentAnalysis
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
