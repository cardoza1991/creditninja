package services

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateLetter(prompt string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4o-mini",
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a credit repair dispute letter generator complying with FCRA."},
				{Role: "user", Content: prompt},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
