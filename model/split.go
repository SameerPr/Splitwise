package model

type Split struct {
	UserId int     `json:"id"`
	Amount float64 `json:"amount"`
}

func NewSplit(userId int, amount float64) *Split {
	return &Split{UserId: userId, Amount: amount}
}

type PercentSplit struct {
	UserId     int     `json:"id"`
	Percentage float64 `json:"percentage"`
}

func NewPercentSplit(userId int, percentage float64) *PercentSplit {
	return &PercentSplit{UserId: userId, Percentage: percentage}
}
