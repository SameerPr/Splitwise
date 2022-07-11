package controller

import (
	"Splitwise/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllBalance() []model.Balance {
	panic("implement me")
}

func (m *MockUserService) DeleteUser(userId int) error {
	return nil
}

func (m *MockUserService) GetUserBalance(userId int) (float64, error) {
	return 0, nil
}

func (m *MockUserService) GetUser(userId int) (*model.User, error) {
	return model.NewUser(1, "Test", "dummy@test.com"), nil
}

func (m *MockUserService) AddUser(user model.User) error {
	return nil
}

type Suite struct {
	suite.Suite
	user            *model.User
	userServiceMock MockUserService
	userController  UserController
}

func (s *Suite) SetupSuite() {
	s.userServiceMock = MockUserService{}
	s.userController = NewUserController(&s.userServiceMock)
	s.user = model.NewUser(1, "Test", "dummy@test.com")
}

func (s *Suite) AfterTest(_, _ string) {
	s.userServiceMock.AssertExpectations(s.T())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Add_User() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/users", s.userController.AddUser)

	jsonValue, _ := json.Marshal(s.user)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodPost, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusCreated, w.Code, "HTTP request status code error")

	expected := "User added successfully."
	assertion.Contains(w.Body.String(), expected)
}

func (s *Suite) Test_Get_User() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users/:id", s.userController.GetUser)

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodGet, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	expected := "{\"Id\":1,\"Name\":\"Test\",\"Email\":\"dummy@test.com\"}"
	assertion.Contains(w.Body.String(), expected)
}

func (s *Suite) Test_Delete_User() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/users/:id", s.userController.DeleteUser)

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	expected := "User deleted successfully."
	assertion.Contains(w.Body.String(), expected)
}

func (s *Suite) Test_Get_User_Balance() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users/balance/:id", s.userController.GetUserBalance)

	req, _ := http.NewRequest(http.MethodGet, "/users/balance/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var assertion = assert.New(s.T())
	assertion.Equal(http.MethodGet, req.Method, "HTTP request method error")
	assertion.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	expected := "Test gets back 0.000000 in total"
	assertion.Contains(w.Body.String(), expected)
}
