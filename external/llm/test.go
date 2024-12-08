package llm

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"neko-acm/prompt"
)

func Test() error {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		//fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return nil
}

func Info() (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt.ModelInfo,
				},
			},
		},
	)
	if err != nil {
		return "", errors.New("ChatCompletion error: " + err.Error())
	}

	return resp.Choices[0].Message.Content, nil
}
