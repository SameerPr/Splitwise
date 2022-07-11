package repository

import (
	"Splitwise/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
	user           *model.User
	userRepository UserRepository
}

func (s *Suite) SetupSuite() {
	s.userRepository = NewUserRepository()
	s.user = model.NewUser(1, "Test", "dummy@test.com")
}

func (s *Suite) AfterTest(_, _ string) {
	s.userRepository = NewUserRepository()
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Add_User() {

	err := s.userRepository.Add(*s.user)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), 1, len(s.userRepository.GetAll()))
}

func (s *Suite) Test_Get_User() {
	err := s.userRepository.Add(*s.user)
	assert.NoError(s.T(), err)

	user, err := s.userRepository.Get(1)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), user, s.user)

}

func (s *Suite) Test_Update_Balance() {
	err := s.userRepository.Add(*s.user)
	assert.NoError(s.T(), err)
	s.userRepository.UpdateBalance(s.user.Id, 100)

	bal, err := s.userRepository.GetBalance(s.user.Id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), bal, float64(100))
}
