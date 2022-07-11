package model

type User struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func NewUser(userId int, name string, email string) *User {
	return &User{Id: userId, Name: name, Email: email}
}

type Balance struct {
	Name   string
	Amount float64
}
