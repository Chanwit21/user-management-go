package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
    "user-management-go/model"
    "user-management-go/service"
)

type UserController struct {
	Service *service.UserService
}

// Validation
func validate(u *model.User) error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

func (ctl *UserController) GetUsers(c *fiber.Ctx) error {
	return c.JSON(ctl.Service.GetAll())
}

func (ctl *UserController) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("userId"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	user, ok := ctl.Service.GetByID(id)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

func (ctl *UserController) CreateUser(c *fiber.Ctx) error {
	var u model.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}
	if err := validate(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	created := ctl.Service.Create(u)
	return c.Status(201).JSON(created)
}

func (ctl *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("userId"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	var u model.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}
	if err := validate(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	updated, err := ctl.Service.Update(id, u)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(updated)
}

func (ctl *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("userId"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	if !ctl.Service.Delete(id) {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.SendStatus(204)
}
