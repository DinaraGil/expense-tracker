package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/DinaraGil/expense-tracker/internal/model"
	"github.com/DinaraGil/expense-tracker/internal/storage"
)

func ListExpenses() (string, error) {
	path, err := storage.GetStoragePath()
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
	defer file.Close()
	scanner := bufio.NewScanner(file)
	idx := 1
	var result string
	for scanner.Scan() {
		// `scanner.Text()` возвращает текущий токен как строку
		j := scanner.Text()
		t := model.Expense{}
		err := json.Unmarshal([]byte(j), &t)
		if errors.Is(err, io.EOF) {
			return result, nil
		}
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf("ID[%d] Date %s Description %s Amount %s\n", idx, t.Date, t.Description, t.Amount)
		idx++
	}
	return result, nil
}
