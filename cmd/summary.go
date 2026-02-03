package cmd

import (
	"fmt"
	"log"

	"github.com/DinaraGil/expense-tracker/internal/service"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("summary called")
		amount, err := service.SummaryExpenses()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Total expenses: $", amount)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
