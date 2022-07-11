package model

import "fmt"

type Expense struct {
	Id           int            `json:"id"`
	Amount       float64        `json:"amount"`
	PaidBy       int            `json:"paidBy"`
	Type         string         `json:"type"`
	Users        []int          `json:"users"`
	Split        []Split        `json:"split"`
	PercentSplit []PercentSplit `json:"percentSplit"`
}

func (expense *Expense) ValidateExactExpense() error {
	var amount float64
	for _, split := range expense.Split {
		amount += split.Amount
	}
	if amount == expense.Amount {
		return nil
	}
	return fmt.Errorf("the amounts do not add up to total %f", expense.Amount)
}

func (expense *Expense) ValidatePercentExpense() error {
	var percent float64
	for _, split := range expense.PercentSplit {
		percent += split.Percentage
	}
	if percent == 100 {
		return nil
	}
	return fmt.Errorf("the percents do not add up to 100")
}
