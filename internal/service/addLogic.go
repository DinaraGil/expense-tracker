package service

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/DinaraGil/expense-tracker/internal/model"
	"github.com/DinaraGil/expense-tracker/internal/storage"
)

func AddToFile(description string, amount float64) (string, error) {
	date := time.Now().Format("2006-01-02")
	expense := model.Expense{Date: date, Description: description, Amount: strconv.FormatFloat(amount, 'f', 2, 64)}
	jsonData, err := json.Marshal(expense)
	if err != nil {
		return "", err
	}
	jsonData = append(jsonData, "\n"...)
	path, err := storage.GetStoragePath()
	if err != nil {
		return "", err
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err = file.Write(jsonData); err != nil {
		return "", err
	}
	return "Expense added successfully", nil
}
