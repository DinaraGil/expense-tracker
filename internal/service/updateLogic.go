package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DinaraGil/expense-tracker/internal/model"
	"github.com/DinaraGil/expense-tracker/internal/storage"
)

func DeleteTempFile() error {
	tempPath, err := storage.GetStoragePath(storage.TempFile)
	if err != nil {
		return err
	}
	if _, err := os.Stat(tempPath); errors.Is(err, os.ErrNotExist) {
		return err
	}
	err = os.Remove(tempPath)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLogic(id int, newDescription string, newAmount float64) (string, error) {
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
	idx := 1
	findExpense := false
	for scanner.Scan() {
		// `scanner.Text()` возвращает текущий токен как строку
		jsonLine := scanner.Text()
		expenseStruct := model.Expense{}
		err := json.Unmarshal([]byte(jsonLine), &expenseStruct)
		//if errors.Is(err, io.EOF) {
		//	return result, nil
		//}
		if err != nil {
			return "", err
		}
		if idx == id {
			findExpense = true
			if newDescription != "" {
				expenseStruct.Description = newDescription
			}
			if newAmount != 0 {
				expenseStruct.Amount = strconv.FormatFloat(newAmount, 'f', 2, 64)
			}
			expenseStruct.Date = time.Now().Format("2006-01-02")
		}
		_, err = AddExpenseToFile(storage.TempFile, &expenseStruct)
		if err != nil {
			_ = DeleteTempFile()
			return "", fmt.Errorf("%w: %w", err)
		}
		idx++
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
