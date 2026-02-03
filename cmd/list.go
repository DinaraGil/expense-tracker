package cmd

import (
	"fmt"
	"log"

	"github.com/DinaraGil/expense-tracker/internal/service"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		msg, err := service.ListExpenses()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
