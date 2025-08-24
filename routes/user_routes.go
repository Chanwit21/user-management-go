package routes

import "github.com/gofiber/fiber/v2"

// Interface for controller methods
type UserControllerInterface interface {
	GetUsers(*fiber.Ctx) error
	GetUser(*fiber.Ctx) error
	CreateUser(*fiber.Ctx) error
	UpdateUser(*fiber.Ctx) error
	DeleteUser(*fiber.Ctx) error
}

// Accept the interface instead of concrete type
func RegisterUserRoutes(app *fiber.App, ctl UserControllerInterface) {
	app.Get("/users", ctl.GetUsers)
	app.Get("/users/:userId", ctl.GetUser)
	app.Post("/users", ctl.CreateUser)
	app.Put("/users/:userId", ctl.UpdateUser)
	app.Delete("/users/:userId", ctl.DeleteUser)
}
