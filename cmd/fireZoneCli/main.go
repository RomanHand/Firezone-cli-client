// main.go
package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"romanhand.ru/fireZoneCli/internal/commands"
	"github.com/joho/godotenv"

)

var (
	apiURL   string
	apiToken string
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получение переменных окружения
	apiURL := os.Getenv("FIREZONE_API_URL")
	apiToken := os.Getenv("FIREZONE_API_TOKEN")
	var rootCmd = &cobra.Command{
		Use:   "firezone-cli",
		Short: "CLI tool to manage Firezone users and connections",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			
		},
	}
	

	rootCmd.AddCommand(
		commands.GetUsersCmd(apiURL, apiToken),
		commands.GetUserCmd(apiURL, apiToken),
		commands.CreateUserCmd(apiURL, apiToken),
		commands.UpdateUserCmd(apiURL, apiToken),
		commands.GetUserCmd(apiURL, apiToken),
		commands.DeleteUserCmd(apiURL, apiToken),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
