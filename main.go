package main

import (
	userController "clean-api/controllers/user"
	"clean-api/database/user"
	"clean-api/pkg/db"
	"github.com/labstack/echo/v4"
)

func main() {

	database := db.InitDB()
	if database == nil {
		panic("Database connection failed")
	}

	instance := echo.New()

	userRepo := user.NewRepository(database)
	UserApiController := userController.NewController(userRepo)
	userController.NewHandlers(instance, UserApiController)

	err := instance.Start(":8080")
	if err != nil {
		return
	}
}
