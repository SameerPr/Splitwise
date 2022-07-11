package service

import (
	"Splitwise/model"
	"Splitwise/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) *userService {
	return &userService{userRepository: ur}
}

type UserService interface {
	AddUser(user model.User) error
	GetUser(userId int) (*model.User, error)
	DeleteUser(userId int) error
	GetUserBalance(userId int) (float64, error)
	GetAllBalance() []model.Balance
}

func (u userService) AddUser(user model.User) error {
	return u.userRepository.Add(user)
}

func (u userService) GetUser(userId int) (*model.User, error) {
	return u.userRepository.Get(userId)
}

func (u userService) DeleteUser(userId int) error {
	return u.userRepository.Delete(userId)
}

func (u userService) GetUserBalance(userId int) (float64, error) {
	return u.userRepository.GetBalance(userId)
}

func (u userService) GetAllBalance() []model.Balance {
	var balances []model.Balance
	balanceSheet := u.userRepository.GetAllBalance()
	for id, amount := range balanceSheet {
		users, _ := u.userRepository.Get(id)
		bal := model.Balance{Name: users.Name, Amount: amount}
		balances = append(balances, bal)
	}
	return balances
}
