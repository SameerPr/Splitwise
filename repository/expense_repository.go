package repository

import (
	"Splitwise/model"
	"fmt"
)

type expenseRepository struct {
	expenseMap map[int]model.Expense
}

func NewExpenseRepository() *expenseRepository {
	return &expenseRepository{
		expenseMap: map[int]model.Expense{},
	}
}

type ExpenseRepository interface {
	AddExpense(expense model.Expense) error
	Get(id int) (*model.Expense, error)
	Delete(id int) error
}

func (e expenseRepository) AddExpense(expense model.Expense) error {
	_, ok := e.expenseMap[expense.Id]
	if ok {
		return fmt.Errorf("expense %d already exists", expense.Id)
	}
	e.expenseMap[expense.Id] = expense
	return nil
}

func (e expenseRepository) Get(id int) (*model.Expense, error) {
	expense, ok := e.expenseMap[id]
	if !ok {
		return nil, fmt.Errorf("expense %d does not exists", id)
	}
	return &expense, nil
}

func (e expenseRepository) Delete(id int) error {
	_, ok := e.expenseMap[id]
	if !ok {
		return fmt.Errorf("expense %d does not exists", id)
	}
	delete(e.expenseMap, id)
	return nil
}
