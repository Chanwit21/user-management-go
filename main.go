package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"user-management-go/controller"
	"user-management-go/routes"
	"user-management-go/service"
)

func main() {
	app := fiber.New()

	// Initialize service & controller
	userService := service.NewUserService()
	userController := &controller.UserController{Service: userService}

	// Register routes
	routes.RegisterUserRoutes(app, userController)

	// Start
	port := "8080"
	fmt.Println("Listening on :" + port)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
