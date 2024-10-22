package llm

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	client := openai.NewClient(option.WithAPIKey(apiKey))

	return &OpenAIClient{client: client}
}

func (c *OpenAIClient) GenerateProblem(
	instructions,
	topicName,
	itemName string,
	generatedProblems []string,
) (string, error) {
	chatCompletion, err := c.client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model: openai.F(openai.ChatModelGPT4o),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf("Instructions: %s", instructions)),
			openai.UserMessage(fmt.Sprintf("Topic: %s", topicName)),
			openai.UserMessage(fmt.Sprintf("Item: %s", itemName)),
			openai.UserMessage(fmt.Sprintf("Past exercises: %v", generatedProblems)),
		}),
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate problem: %v", err)
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
