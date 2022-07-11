package service

import (
	"Splitwise/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) GetAllBalance() map[int]float64 {
	//TODO implement me
	panic("implement me")
}

func (m *userRepositoryMock) Get(i int) (*model.User, error) {
	panic("implement me")
}

func (m *userRepositoryMock) Delete(i int) error {
	panic("implement me")
}

func (m *userRepositoryMock) GetBalance(i int) (float64, error) {
	panic("implement me")
}

func (m *userRepositoryMock) UpdateBalance(id int, amount float64) error {
	panic("implement me")
}

func (m *userRepositoryMock) GetAll() map[int]model.User {
	panic("implement me")
}

func (m *userRepositoryMock) Add(user model.User) error {
	return nil
}

type Suite struct {
	suite.Suite
	user               *model.User
	userRepositoryMock userRepositoryMock
	userService        UserService
}

func (s *Suite) SetupSuite() {
	s.userRepositoryMock = userRepositoryMock{}
	s.userService = NewUserService(&s.userRepositoryMock)
	s.user = model.NewUser(1, "Test", "dummy@test.com")
}

func (s *Suite) AfterTest(_, _ string) {
	s.userRepositoryMock.AssertExpectations(s.T())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Add_User() {
	s.userRepositoryMock.On("AddUser", s.user).Return(nil)

	err := s.userService.AddUser(*s.user)
	assert.NoError(s.T(), err)

	s.userRepositoryMock.MethodCalled("AddUser", s.user)
}
