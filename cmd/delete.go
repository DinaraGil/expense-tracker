package cmd

import (
	"fmt"
	"log"

	"github.com/DinaraGil/expense-tracker/internal/service"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			log.Fatal(err)
		}
		if id <= 0 {
			log.Fatal("id is required, should be positive integer")
		}

		msg, err := service.DeleteLogic(id)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(msg)
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntP(
		"id",
		"i",
		0,
		"Id of the expense",
	)
}
