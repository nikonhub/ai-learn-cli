package application

import (
	"context"
	"log"

	"github.com/nikonhub/ai-learn-cli/ent/settings"
	"github.com/nikonhub/ai-learn-cli/internal/database"
	"github.com/nikonhub/ai-learn-cli/internal/llm"
	"github.com/nikonhub/ai-learn-cli/internal/topic"
)

type Application struct {
	db           *database.DB
	llmClient    llm.Client
	TopicService *topic.TopicService
}

func NewApplication() *Application {
	db := database.NewDB()
	openaiAPIKey, err := db.Client.Settings.Query().Where(settings.Key("OPENAI_API_KEY")).First(context.Background())
	if err != nil {
		log.Fatalf("failed to query settings: %v", err)
	}

	llmClient := llm.NewOpenAIClient(openaiAPIKey.Value)

	topicService := topic.NewTopicService(db, llmClient)

	return &Application{db: db, llmClient: llmClient, TopicService: topicService}
}
