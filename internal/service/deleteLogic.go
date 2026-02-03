package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DinaraGil/expense-tracker/internal/model"
	"github.com/DinaraGil/expense-tracker/internal/storage"
)

func DeleteLogic(id int) (string, error) {
	path, err := storage.GetStoragePath(storage.ConstFile)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	idx := 0
	findExpense := false
	for scanner.Scan() {
		idx++
		jsonLine := scanner.Text()
		expenseStruct := model.Expense{}
		err := json.Unmarshal([]byte(jsonLine), &expenseStruct)

		if err != nil {
			return "", err
		}
		if idx == id {
			findExpense = true
			continue
		}
		_, err = AddExpenseToFile(storage.TempFile, &expenseStruct)
		if err != nil {
			_ = DeleteTempFile()
			return "", fmt.Errorf("%w: %w", err)
		}
	}
	if !findExpense {
		_ = DeleteTempFile()
		return "", fmt.Errorf("expense with id %d doesn't exist", id)
	}
	err = file.Close()
	if err != nil {
		_ = DeleteTempFile()
		return "", err
	}
	tempPath, err := storage.GetStoragePath(storage.TempFile)
	if err != nil {
		_ = DeleteTempFile()
		return "", err
	}
	constFile, err := storage.GetStoragePath(storage.ConstFile)
	if err != nil {
		_ = DeleteTempFile()
		return "", err
	}
	err = os.Rename(tempPath, constFile)
	if err != nil {
		_ = DeleteTempFile()
		return "", err
	}
	_ = DeleteTempFile()

	return "Expense updates successfully", nil
}
