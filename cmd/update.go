/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DinaraGil/expense-tracker/internal/service"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			log.Fatal(err)
		}
		if id <= 0 {
			log.Fatal("id is required, should be positive integer")
		}

		var description string
		description, err = cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}

		var amount float64
		amount, err = cmd.Flags().GetFloat64("amount")
		if err != nil {
			log.Fatal(err)
		}
		if amount < 0 {
			log.Fatal("amount should be positive")
		}
		if amount == 0 && description == "" {
			log.Fatal("amount or description required")
		}
		msg, err := service.UpdateLogic(id, description, amount)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntP(
		"id",
		"i",
		0,
		"Id of the expense",
	)
	updateCmd.Flags().StringP(
		"description",
		"d",
		"",
		"New description of the expense",
	)
	updateCmd.Flags().Float64P(
		"amount",
		"a",
		0,
		"New amount of the expense",
	)
}
