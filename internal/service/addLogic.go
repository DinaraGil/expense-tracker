package service

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/DinaraGil/expense-tracker/internal/model"
	"github.com/DinaraGil/expense-tracker/internal/storage"
)

func AddExpenseToFile(filename storage.Filename, expense *model.Expense) (string, error) {
	jsonData, err := json.Marshal(expense)
	if err != nil {
		return "", err
	}
	jsonData = append(jsonData, "\n"...)
	path, err := storage.GetStoragePath(filename)
	if err != nil {
		return "", err
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	if _, err = file.Write(jsonData); err != nil {
		return "", err
	}
	err = file.Close()
	if err != nil {
		return "", err
	}
	return "Expense added successfully", nil
}
func AddLogic(description string, amount float64) (string, error) {
	date := time.Now().Format("2006-01-02")
	expense := model.Expense{Date: date, Description: description, Amount: strconv.FormatFloat(amount, 'f', 2, 64)}
	msg, err := AddExpenseToFile(storage.ConstFile, &expense)
	return msg, err
}
