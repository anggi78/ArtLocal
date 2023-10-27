package services

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type evtRecService struct {}

type EvtInterface interface {
	Reccomend(userInput, category, location, openAIKey string) (string, error)
}

func NewReccomEvent() EvtInterface {
	return &evtRecService{}
}

func (s *evtRecService) Reccomend(userInput, category, location, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: "Halo, saya Artlo akan membantu mu menemukan event apa yang ingin kamu cari",
		},
		{
			Role: openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := s.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (s *evtRecService) getCompletionFromMessages (
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: model,
			Messages: messages,
		},
	)
	return resp, err
}