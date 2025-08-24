package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"user-management-go/model"
	"user-management-go/service"

	"github.com/gofiber/fiber/v2"
)

func setupApp() (*fiber.App, *UserController) {
	app := fiber.New()
	svc := service.NewUserService()
	ctl := &UserController{Service: svc}

	app.Get("/users", ctl.GetUsers)
	app.Get("/users/:userId", ctl.GetUser)
	app.Post("/users", ctl.CreateUser)
	app.Put("/users/:userId", ctl.UpdateUser)
	app.Delete("/users/:userId", ctl.DeleteUser)

	return app, ctl
}

func TestCreateUser(t *testing.T) {
	app, _ := setupApp()

	user := model.User{
		Name:     "John Doe",
		Username: "johndoe",
		Email:    "john@example.com",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Errorf("expected status 201, got %d", resp.StatusCode)
	}
}

func TestGetUsers(t *testing.T) {
	app, _ := setupApp()
	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestGetUser(t *testing.T) {
	app, ctl := setupApp()
	created := ctl.Service.Create(model.User{Name: "Jane", Username: "jane", Email: "jane@example.com"})

	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", created.ID), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	app, ctl := setupApp()
	created := ctl.Service.Create(model.User{Name: "Alice", Username: "alice", Email: "alice@example.com"})

	updatedUser := model.User{Name: "Alice Updated", Username: "alice", Email: "alice@example.com"}
	body, _ := json.Marshal(updatedUser)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/users/%d", created.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	app, ctl := setupApp()
	created := ctl.Service.Create(model.User{Name: "Bob", Username: "bob", Email: "bob@example.com"})

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%d", created.ID), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	if resp.StatusCode != 204 {
		t.Errorf("expected status 204, got %d", resp.StatusCode)
	}
}
