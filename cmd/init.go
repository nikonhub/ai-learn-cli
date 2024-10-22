package cmd

import (
	"log"

	"github.com/nikonhub/ai-learn-cli/internal/database"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database schema",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDB()

		if err := db.InitSchema(); err != nil {
			log.Fatalf("failed to initialize database schema: %v", err)
		}

		cmd.Println("Database schema initialized successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
