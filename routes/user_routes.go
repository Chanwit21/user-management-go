package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-management-go/controller"
)

func RegisterUserRoutes(app *fiber.App, ctl *controller.UserController) {
	app.Get("/users", ctl.GetUsers)
	app.Get("/users/:userId", ctl.GetUser)
	app.Post("/users", ctl.CreateUser)
	app.Put("/users/:userId", ctl.UpdateUser)
	app.Delete("/users/:userId", ctl.DeleteUser)
}
