package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/DinaraGil/expense-tracker/internal/model"
	"github.com/DinaraGil/expense-tracker/internal/storage"
)

func SummaryExpenses() (float64, error) {
	path, err := storage.GetStoragePath(storage.ConstFile)
	if err != nil {
		return 0, err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return 0, err
	}
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result float64
	for scanner.Scan() {
		// `scanner.Text()` возвращает текущий токен как строку
		jsonLine := scanner.Text()
		expenseStruct := model.Expense{}
		err := json.Unmarshal([]byte(jsonLine), &expenseStruct)
		if errors.Is(err, io.EOF) {
			return result, nil
		}
		if err != nil {
			return 0, err
		}
		floatAmount, _ := strconv.ParseFloat(expenseStruct.Amount, 64)
		result += floatAmount
	}
	return result, nil
}
