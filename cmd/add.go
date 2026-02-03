package cmd

import (
	"fmt"
	"log"

	"github.com/DinaraGil/expense-tracker/internal/service"
	"github.com/spf13/cobra"
)

// addExpenseCmd represents the addExpense command
var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}
		if description == "" {
			log.Fatal("description is required")
		}
		amount, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			log.Fatal(err)
		}
		if amount <= 0 {
			log.Fatal("amount should be positive")
		}
		msg, err := service.AddLogic(description, amount)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(addExpenseCmd)
	addExpenseCmd.Flags().StringP(
		"description",
		"d",
		"",
		"Description of the expense",
	)
	addExpenseCmd.Flags().Float64P(
		"amount",
		"a",
		0,
		"Amount of the expense",
	)
}
