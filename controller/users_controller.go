package controller

import (
	"Splitwise/model"
	"Splitwise/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetAllBalance(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetUserBalance(ctx *gin.Context)
}

func NewUserController(s service.UserService) UserController {
	return userController{
		userService: s,
	}
}

type userController struct {
	userService service.UserService
}

func (u userController) AddUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := u.userService.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Adding users in database failed. "+err.Error())
		return
	}
	c.JSON(http.StatusCreated, "User added successfully.")
}

func (u userController) GetUser(c *gin.Context) {
	pathParam := c.Param("id")
	userId, _ := strconv.Atoi(pathParam)

	user, err := u.userService.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "User not found "+err.Error())
		return
	}
	jsonData, _ := json.Marshal(&user)
	c.Data(http.StatusOK, "user/json", jsonData)
}

func (u userController) DeleteUser(c *gin.Context) {
	pathParam := c.Param("id")
	userId, _ := strconv.Atoi(pathParam)

	err := u.userService.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "User not found "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "User deleted successfully.")
}

func (u userController) GetUserBalance(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)
	user, err := u.userService.GetUser(id)
	balance, err := u.userService.GetUserBalance(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "User not found "+err.Error())
		return
	}
	if balance >= 0 {
		c.JSON(http.StatusOK, fmt.Sprintf("%s gets back %f in total", user.Name, balance))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("%s owes %f in total", user.Name, -1*balance))
	}
}

func (u userController) GetAllBalance(c *gin.Context) {
	balances := u.userService.GetAllBalance()

	var response []string
	for _, balance := range balances {
		if balance.Amount >= 0 {
			response = append(response, fmt.Sprintf("%s gets back %f in total", balance.Name, balance.Amount))
		} else {
			response = append(response, fmt.Sprintf("%s owes %f in total", balance.Name, -1*balance.Amount))
		}
	}
	c.JSON(http.StatusOK, response)
}
