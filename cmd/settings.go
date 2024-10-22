package cmd

import (
	"context"
	"log"

	"github.com/nikonhub/ai-learn-cli/internal/database"
	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Manage settings",
	Long:  "Command for managing settings.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var addSettingCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a new setting",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		db := database.NewDB()

		db.Client.Settings.Create().SetKey(key).SetValue(value).SaveX(context.Background())
	},
}

var listSettingsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all settings",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDB()

		settings, err := db.Client.Settings.Query().All(context.Background())
		if err != nil {
			log.Fatalf("failed to query settings: %v", err)
		}

		for _, setting := range settings {
			cmd.Println(setting.Key, setting.Value)
		}
	},
}

var availableKeysCmd = &cobra.Command{
	Use:   "available-keys",
	Short: "List available settings keys",
	Run: func(cmd *cobra.Command, args []string) {
		availableKeys := []struct {
			Key         string
			Description string
		}{
			{"OPENAI_API_KEY", "API key for OpenAI service"},
		}

		for _, key := range availableKeys {
			cmd.Printf("%-20s %s\n", key.Key, key.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(settingsCmd)
	settingsCmd.AddCommand(addSettingCmd)
	settingsCmd.AddCommand(listSettingsCmd)
	settingsCmd.AddCommand(availableKeysCmd)
}
