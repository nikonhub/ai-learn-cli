package cmd

import (
	"log"
	"strconv"

	"github.com/nikonhub/ai-learn-cli/internal/application"
	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Manage topics",
	Long:  "Command for creating, listing, and managing learning topics.",
}

var addTopicCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topicName := args[0]

		app := application.NewApplication()
		topic := app.TopicService.CreateTopic(topicName, "")
		cmd.Printf("Topic created: %d : %s\n", topic.ID, topic.Name)
	},
}

var listTopicsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all topics",
	Run: func(cmd *cobra.Command, args []string) {
		app := application.NewApplication()
		topics, err := app.TopicService.ListTopics()
		if err != nil {
			log.Fatalf("failed to query topics: %v", err)
		}

		for _, topic := range topics {
			cmd.Printf("%d : %s\n", topic.ID, topic.Name)
		}
	},
}

var learnTopicCmd = &cobra.Command{
	Use:   "learn [topic-id]",
	Short: "Learn a topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topicID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to parse topic id: %v", err)
		}

		app := application.NewApplication()
		item, generatedProblem, err := app.TopicService.Learn(topicID)
		if err != nil {
			log.Fatalf("failed to learn topic: %v", err)
		}

		cmd.Printf("Item: %d : %s\nProblem: %s\n", item.ID, item.Name, generatedProblem.ProblemText)
	},
}

func init() {
	rootCmd.AddCommand(topicCmd)
	topicCmd.AddCommand(addTopicCmd)
	topicCmd.AddCommand(listTopicsCmd)
	topicCmd.AddCommand(learnTopicCmd)
}
