package model

type Expense struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
}
