package controller

import (
	"Splitwise/model"
	"Splitwise/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExpenseController interface {
	AddExpense(ctx *gin.Context)
	DeleteExpense(ctx *gin.Context)
}

func NewExpenseController(s service.ExpenseService, u service.UserService) ExpenseController {
	return expenseController{
		expenseService: s,
		userService:    u,
	}
}

type expenseController struct {
	expenseService service.ExpenseService
	userService    service.UserService
}

func (u expenseController) AddExpense(c *gin.Context) {
	var expense model.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, "error: "+err.Error())
		return
	}
	var err error
	err = u.expenseService.AddExpense(expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Adding expense in database failed. "+err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Successfully added expense.")
}

func (u expenseController) DeleteExpense(c *gin.Context) {
	pathParam := c.Param("id")
	expenseId, _ := strconv.Atoi(pathParam)

	var err error
	err = u.expenseService.DeleteExpense(expenseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Deleting expense in database failed. "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted expense.")
}
