package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	// important: reset store to known state for tests
	store = newUserStore()
	app := fiber.New()
	app.Get("/users", listUsersHandler)
	app.Get("/users/:userId", getUserHandler)
	app.Post("/users", createUserHandler)
	app.Put("/users/:userId", updateUserHandler)
	app.Delete("/users/:userId", deleteUserHandler)
	return app
}

func TestListUsers(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}
	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(users) < 1 {
		t.Fatalf("expected at least one user, got %d", len(users))
	}
}

func TestCreateUserValidation(t *testing.T) {
	app := setupApp()
	// missing email
	body := map[string]string{
		"name":     "Test User",
		"username": "testuser",
		// "email" missing
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request for missing email, got %d", resp.StatusCode)
	}
}
