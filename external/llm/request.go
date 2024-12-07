package llm

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
)

func RequestMessage(msg openai.ChatCompletionMessage) (openai.ChatCompletionMessage, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    config.Model,
			Messages: []openai.ChatCompletionMessage{msg},
		},
	)

	if err != nil {
		return openai.ChatCompletionMessage{}, errors.New("ChatCompletion error: " + err.Error())
	}

	return resp.Choices[0].Message, nil
}

func RequestMessages(msgs []openai.ChatCompletionMessage) (openai.ChatCompletionMessage, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    config.Model,
			Messages: msgs,
		},
	)

	if err != nil {
		return openai.ChatCompletionMessage{}, errors.New("ChatCompletion error: " + err.Error())
	}

	return resp.Choices[0].Message, nil
}
