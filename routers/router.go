package routers

import (
	"Splitwise/controller"
	"Splitwise/repository"
	"Splitwise/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Route() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router.POST("/users", userController.AddUser)
	router.GET("/users/:id", userController.GetUser)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.GET("/users/balance/:id", userController.GetUserBalance)
	router.GET("/users/balance", userController.GetAllBalance)

	expenseRepo := repository.NewExpenseRepository()
	expenseService := service.NewExpenseService(expenseRepo, userRepo)
	expenseController := controller.NewExpenseController(expenseService, userService)

	router.POST("/expense", expenseController.AddExpense)
	router.DELETE("/expense/:id", expenseController.DeleteExpense)

	log.Println(fmt.Sprintf("listening on port %d", 8080))
	if err := router.Run(fmt.Sprintf(":%d", 8080)); err != nil {
		log.Fatal(err)
	}
}
