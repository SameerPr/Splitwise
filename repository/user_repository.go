package repository

import (
	"Splitwise/model"
	"fmt"
)

type userRepository struct {
	UserMap      map[int]model.User
	BalanceSheet map[int]float64
}

func NewUserRepository() *userRepository {
	userMap := map[int]model.User{}
	balanceSheet := map[int]float64{}
	return &userRepository{UserMap: userMap, BalanceSheet: balanceSheet}
}

type UserRepository interface {
	Add(user model.User) error
	Get(int) (*model.User, error)
	GetAll() map[int]model.User
	Delete(int) error
	GetBalance(int) (float64, error)
	GetAllBalance() map[int]float64
	UpdateBalance(id int, amount float64) error
}

func (u userRepository) Add(user model.User) error {
	_, ok := u.UserMap[user.Id]
	if ok {
		return fmt.Errorf("user %d already exists", user.Id)
	}
	u.UserMap[user.Id] = user
	u.BalanceSheet[user.Id] = float64(0)
	return nil
}

func (u userRepository) Get(id int) (*model.User, error) {
	user, ok := u.UserMap[id]
	if !ok {
		return nil, fmt.Errorf("user %d does not exists", user.Id)
	}
	return &user, nil
}

func (u userRepository) GetAll() map[int]model.User {
	return u.UserMap
}

func (u userRepository) GetBalance(userId int) (float64, error) {
	for id, _ := range u.BalanceSheet {
		if id == userId {
			return u.BalanceSheet[userId], nil
		}
	}
	return 0, fmt.Errorf("user %d does not exists", userId)
}

func (u userRepository) GetAllBalance() map[int]float64 {
	return u.BalanceSheet
}

func (u userRepository) UpdateBalance(userId int, amount float64) error {
	for id, _ := range u.BalanceSheet {
		if id == userId {
			u.BalanceSheet[userId] += amount
			return nil
		}
	}
	return fmt.Errorf("user %d does not exists", userId)
}

func (u userRepository) Delete(id int) error {
	user, ok := u.UserMap[id]
	if !ok {
		return fmt.Errorf("user %d does not exists", user.Id)
	}
	delete(u.UserMap, id)
	return nil
}
