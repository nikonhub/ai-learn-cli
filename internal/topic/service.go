package topic

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/nikonhub/ai-learn-cli/ent"
	"github.com/nikonhub/ai-learn-cli/ent/item"
	"github.com/nikonhub/ai-learn-cli/ent/itemprogress"
	"github.com/nikonhub/ai-learn-cli/ent/topic"
	"github.com/nikonhub/ai-learn-cli/internal/database"
	"github.com/nikonhub/ai-learn-cli/internal/llm"
	"github.com/nikonhub/ai-learn-cli/internal/utils"
)

var defaultInstructions = `
Instructions:
- Generate a daily exercise based on the topic and current item.
- Don't repeat the same exercise. Look at the past exercises and generate a new one.
- The response should be a valid JSON object with the following fields:
  - "problem": The exercise instructions.
  - "solution": The solution to the exercise.
  - "hints": Hints to solve the exercise.
`

type GeneratedProblem struct {
	Problem  string `json:"problem"`
	Solution string `json:"solution"`
	Hints    string `json:"hints"`
}

type TopicService struct {
	db        *database.DB
	llmClient llm.Client
}

func NewTopicService(db *database.DB, llmClient llm.Client) *TopicService {
	return &TopicService{db: db, llmClient: llmClient}
}

func (s *TopicService) CreateTopic(topicName, instructions string) *ent.Topic {
	if instructions == "" {
		instructions = defaultInstructions
	}

	return s.db.Client.Topic.
		Create().
		SetName(topicName).
		SetInstructions(instructions).
		SaveX(context.Background())
}

func (s *TopicService) ListTopics() ([]*ent.Topic, error) {
	return s.db.Client.Topic.Query().All(context.Background())
}

func (s *TopicService) AddItemToTopic(topicID int, itemName string) error {
	topic, err := s.db.Client.Topic.Get(context.Background(), topicID)
	if err != nil {
		return fmt.Errorf("failed to query topic: %v", err)
	}

	item := s.db.Client.Item.
		Create().
		SetTopicID(topic.ID).
		SetName(itemName).
		SaveX(context.Background())

	s.db.Client.ItemProgress.
		Create().
		SetItemID(item.ID).
		SetNextReviewDate(time.Now()).
		SaveX(context.Background())

	return nil
}

func (s *TopicService) RemoveItem(itemID int) {
	s.db.Client.Item.DeleteOneID(itemID).ExecX(context.Background())
}

func (s *TopicService) Learn(topicID int) (*ent.Item, *ent.GeneratedProblem, error) {
	t, err := s.db.Client.Topic.Get(context.Background(), topicID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query topic: %v", err)
	}

	itemProgress, err := s.db.Client.ItemProgress.
		Query().
		Where(
			itemprogress.HasItemWith(
				item.HasTopicWith(
					topic.ID(t.ID),
				),
			),
			itemprogress.NextReviewDateLTE(time.Now()),
		).
		Order(ent.Asc(itemprogress.FieldNextReviewDate)).
		WithItem(func(q *ent.ItemQuery) {
			q.WithGeneratedProblems()
		}).
		First(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query item progress: %v", err)
	}

	item := itemProgress.Edges.Item
	problems := itemProgress.Edges.Item.Edges.GeneratedProblems

	if problems == nil {
		problems = []*ent.GeneratedProblem{}
	}

	problemStrs := utils.Map(problems, func(p *ent.GeneratedProblem) string {
		return p.ProblemText
	})

	if len(problemStrs) == 0 {
		problemStrs = append(problemStrs, "No exercises yet.")
	}

	problemStr, err := s.llmClient.GenerateProblem(t.Instructions, t.Name, item.Name, problemStrs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate problem: %v", err)
	}

	problemJson := GeneratedProblem{}
	err = json.Unmarshal([]byte(problemStr), &problemJson)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal generated problem: %v", err)
	}

	problem := s.db.Client.GeneratedProblem.
		Create().
		SetItemID(item.ID).
		SetProblemText(problemJson.Problem).
		SaveX(context.Background())

	return item, problem, nil
}

func (s *TopicService) FeedbackItem(itemID int, feedback string) error {
	itemProgress, err := s.db.Client.ItemProgress.Get(context.Background(), itemID)
	if err != nil {
		return fmt.Errorf("failed to query item progress: %v", err)
	}

	switch feedback {
	case "easy":
		itemProgress.EaseFactor += 0.15
		itemProgress.IntervalDays = int(float64(itemProgress.IntervalDays) * itemProgress.EaseFactor)
		itemProgress.Streak += 1
	case "medium":
		itemProgress.IntervalDays = int(float64(itemProgress.IntervalDays) * itemProgress.EaseFactor)
		itemProgress.Streak += 1
	case "hard":
		itemProgress.EaseFactor = math.Max(1.3, itemProgress.EaseFactor-0.2)
		itemProgress.IntervalDays = int(math.Max(1, float64(itemProgress.IntervalDays)*itemProgress.EaseFactor))
		itemProgress.Streak = int(math.Max(0, float64(itemProgress.Streak)-1))
	default:
		return fmt.Errorf("invalid feedback: %s", feedback)
	}

	itemProgress.NextReviewDate = time.Now().AddDate(0, 0, itemProgress.IntervalDays)

	_, err = s.db.Client.ItemProgress.UpdateOneID(itemID).
		SetEaseFactor(itemProgress.EaseFactor).
		SetIntervalDays(itemProgress.IntervalDays).
		SetStreak(itemProgress.Streak).
		SetNextReviewDate(itemProgress.NextReviewDate).
		Save(context.Background())

	if err != nil {
		return fmt.Errorf("failed to update item progress: %v", err)
	}

	return nil
}

func (s *TopicService) ListTopicItems(topicID int) ([]*ent.Item, error) {
	topic, err := s.db.Client.Topic.Get(context.Background(), topicID)
	if err != nil {
		return nil, fmt.Errorf("failed to query topic: %v", err)
	}

	return topic.QueryItems().All(context.Background())
}
