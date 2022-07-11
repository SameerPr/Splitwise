package service

import (
	"Splitwise/model"
	"Splitwise/repository"
	"fmt"
)

type expenseService struct {
	userRepository    repository.UserRepository
	expenseRepository repository.ExpenseRepository
}

func NewExpenseService(er repository.ExpenseRepository, ur repository.UserRepository) *expenseService {
	return &expenseService{userRepository: ur, expenseRepository: er}
}

type ExpenseService interface {
	AddExpense(expense model.Expense) error
	AddEqualExpense(expense model.Expense) error
	AddExactExpense(expense model.Expense) error
	AddPercentExpense(expense model.Expense) error
	DeleteExpense(expenseId int) error
}

func (e expenseService) AddExpense(expense model.Expense) error {

	if expense.Type == "Equal" {
		return e.AddEqualExpense(expense)
	} else if expense.Type == "Exact" {
		if err1 := expense.ValidateExactExpense(); err1 != nil {
			return fmt.Errorf(err1.Error())
		}
		return e.AddExactExpense(expense)
	} else if expense.Type == "Percentage" {
		if err1 := expense.ValidatePercentExpense(); err1 != nil {
			return fmt.Errorf(err1.Error())
		}
		return e.AddPercentExpense(expense)
	}
	return nil
}

func (e expenseService) AddEqualExpense(expense model.Expense) error {

	err := e.userRepository.UpdateBalance(expense.PaidBy, expense.Amount)
	if err != nil {
		return err
	}
	individualAmount := expense.Amount / float64(len(expense.Users))
	for _, user := range expense.Users {
		err := e.userRepository.UpdateBalance(user, -1*individualAmount)
		if err != nil {
			return err
		}
	}
	err = e.expenseRepository.AddExpense(expense)
	if err != nil {
		return err
	}
	return nil
}

func (e expenseService) AddExactExpense(expense model.Expense) error {
	err := e.userRepository.UpdateBalance(expense.PaidBy, expense.Amount)
	if err != nil {
		return err
	}
	for _, split := range expense.Split {
		err := e.userRepository.UpdateBalance(split.UserId, -1*split.Amount)
		if err != nil {
			return err
		}
	}
	err = e.expenseRepository.AddExpense(expense)
	if err != nil {
		return err
	}
	return nil
}

func (e expenseService) AddPercentExpense(expense model.Expense) error {
	err := e.userRepository.UpdateBalance(expense.PaidBy, expense.Amount)
	if err != nil {
		return err
	}
	for _, split := range expense.PercentSplit {
		err := e.userRepository.UpdateBalance(split.UserId, -1*expense.Amount*(split.Percentage/100))
		if err != nil {
			return err
		}
	}
	err = e.expenseRepository.AddExpense(expense)
	if err != nil {
		return err
	}
	return nil
}

func (e expenseService) DeleteExpense(id int) error {
	expense, err := e.expenseRepository.Get(id)
	if err != nil {
		return err
	}
	fmt.Println(*expense)
	if expense.Type == "Equal" {
		expense.Amount *= -1
		return e.AddEqualExpense(*expense)
	} else if expense.Type == "Exact" {
		return e.RemoveExactExpense(*expense)
	} else if expense.Type == "Percentage" {
		expense.Amount *= -1
		return e.AddPercentExpense(*expense)
	}
	err = e.expenseRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (e expenseService) RemoveExactExpense(expense model.Expense) error {
	err := e.userRepository.UpdateBalance(expense.PaidBy, -1*expense.Amount)
	if err != nil {
		return err
	}
	for _, split := range expense.Split {
		err := e.userRepository.UpdateBalance(split.UserId, split.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}
