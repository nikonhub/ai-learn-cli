package cmd

import (
	"log"
	"strconv"

	"github.com/nikonhub/ai-learn-cli/internal/application"
	"github.com/spf13/cobra"
)

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Manage items",
	Long:  "Command for creating, listing, and managing learning items.",
}

var addItemCmd = &cobra.Command{
	Use:   "add [topic-id] [item]",
	Short: "Add a new item",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		topicID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to parse topic id: %v", err)
		}

		itemName := args[1]

		app := application.NewApplication()
		if err := app.TopicService.AddItemToTopic(topicID, itemName); err != nil {
			log.Fatalf("failed to create item: %v", err)
		}
	},
}

var listItemsCmd = &cobra.Command{
	Use:   "list [topic-id]",
	Short: "List all items",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topicID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to parse topic id: %v", err)
		}

		app := application.NewApplication()
		items, err := app.TopicService.ListTopicItems(topicID)
		if err != nil {
			log.Fatalf("failed to query items: %v", err)
		}

		for _, item := range items {
			cmd.Printf("%d : %s\n", item.ID, item.Name)
		}
	},
}

var remoteItemCmd = &cobra.Command{
	Use:   "remove [item-id]",
	Short: "Remove an item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		itemID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to parse item id: %v", err)
		}

		app := application.NewApplication()
		app.TopicService.RemoveItem(itemID)
	},
}

var feedbackItemCmd = &cobra.Command{
	Use:   "feedback [item-id] [feedback]",
	Short: "Feedback on an item : easy, medium, hard",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		itemID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to parse item id: %v", err)
		}

		feedback := args[1]

		app := application.NewApplication()
		if err := app.TopicService.FeedbackItem(itemID, feedback); err != nil {
			log.Fatalf("failed to feedback item: %v", err)
		}

		cmd.Println("Item feedback recorded")
	},
}

func init() {
	rootCmd.AddCommand(itemCmd)
	itemCmd.AddCommand(addItemCmd)
	itemCmd.AddCommand(listItemsCmd)
	itemCmd.AddCommand(feedbackItemCmd)
	itemCmd.AddCommand(remoteItemCmd)
}
