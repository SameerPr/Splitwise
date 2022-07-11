package controller

import (
	"Splitwise/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockExpenseService struct {
	mock.Mock
}

func (m *MockExpenseService) DeleteExpense(expenseId int) error {
	return nil
}

func (m *MockExpenseService) AddEqualExpense(expense model.Expense) error {
	panic("implement me")
}

func (m *MockExpenseService) AddExactExpense(expense model.Expense) error {
	return expense.ValidateExactExpense()
}

func (m *MockExpenseService) AddPercentExpense(expense model.Expense) error {
	return expense.ValidatePercentExpense()
}

func (m *MockExpenseService) AddExpense(expense model.Expense) error {
	if expense.Type == "Equal" {
		return m.AddEqualExpense(expense)
	} else if expense.Type == "Exact" {
		return m.AddExactExpense(expense)
	} else if expense.Type == "Percentage" {
		return m.AddPercentExpense(expense)
	}
	return nil
}

type ExpenseSuite struct {
	suite.Suite
	user               *model.User
	expense            *model.Expense
	expenseServiceMock MockExpenseService
	userServiceMock    MockUserService
	expenseController  ExpenseController
}

func (s *ExpenseSuite) SetupSuite() {
	s.expenseServiceMock = MockExpenseService{}
	s.expenseController = NewExpenseController(&s.expenseServiceMock, &s.userServiceMock)
	s.user = model.NewUser(1, "Test", "dummy@test.com")
}

func (s *ExpenseSuite) AfterTest(_, _ string) {
	s.userServiceMock.AssertExpectations(s.T())
}

func TestInitExpense(t *testing.T) {
	suite.Run(t, new(ExpenseSuite))
}

func (s *ExpenseSuite) Test_Add_Exact_Expense() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/expense", s.expenseController.AddExpense)

	data := `{
    "id": 1,
    "amount": 500,
    "paidBy": 2,
	"type": "Exact",
    "split": [
        {
            "id": 1,
            "amount": 200
        },
        {
            "id": 2,
            "amount": 300
        }
    ]
}`
	req, _ := http.NewRequest(http.MethodPost, "/expense", ioutil.NopCloser(bytes.NewReader([]byte(data))))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodPost, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusCreated, w.Code, "HTTP request status code error")

	expected := "Successfully added expense."
	assertion.Contains(w.Body.String(), expected)
}

func (s *ExpenseSuite) Test_Add_Exact_Expense_Unequal_Amount() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/expense", s.expenseController.AddExpense)

	data := `{
    "id": 1,
    "amount": 500,
    "paidBy": 2,
	"type": "Exact",
    "split": [
        {
            "id": 1,
            "amount": 200
        },
        {
            "id": 2,
            "amount": 200
        }
    ]
}`
	req, _ := http.NewRequest(http.MethodPost, "/expense", ioutil.NopCloser(bytes.NewReader([]byte(data))))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodPost, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusBadRequest, w.Code, "HTTP request status code error")

	expected := "Adding expense in database failed. the amounts do not add up to total 500.000000"
	assertion.Contains(w.Body.String(), expected)
}

func (s *ExpenseSuite) Test_Add_Percent_Expense() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/expense", s.expenseController.AddExpense)

	data := `{
    "id": 1,
    "amount": 500,
    "paidBy": 1,
	"type": "Percentage",
    "percentSplit": [
        {
            "id": 1,
            "percentage": 30
        },
        {
            "id": 2,
            "percentage": 70
        }
    ]
}`
	req, _ := http.NewRequest(http.MethodPost, "/expense", ioutil.NopCloser(bytes.NewReader([]byte(data))))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodPost, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusCreated, w.Code, "HTTP request status code error")

	expected := "Successfully added expense."
	assertion.Contains(w.Body.String(), expected)
}

func (s *ExpenseSuite) Test_Add_Percent_Expense_Unequal_Percent() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/expense/:id", s.expenseController.DeleteExpense)

	req, _ := http.NewRequest(http.MethodDelete, "/expense/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	expected := "Successfully deleted expense."
	assertion.Contains(w.Body.String(), expected)
}
