package cmd

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	Date        string
	Description string
	Amount      string
}

func addToFile(description string, amount float64) (string, error) {
	date := time.Now().Format("2006-01-02")
	expense := Expense{Date: date, Description: description, Amount: strconv.FormatFloat(amount, 'f', 2, 64)}
	jsonData, err := json.MarshalIndent(expense, "", " ")
	if err != nil {
		return "", err
	}
	file, err := os.OpenFile("./internal/tasks.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err = file.Write(jsonData); err != nil {
		return "", err
	}
	return "Expense added successfully", nil
}
