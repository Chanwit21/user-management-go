package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Mock controller implementing the interface
type MockUserController struct{}

func (m *MockUserController) GetUsers(c *fiber.Ctx) error   { return c.Status(200).SendString("GetUsers") }
func (m *MockUserController) GetUser(c *fiber.Ctx) error    { return c.Status(200).SendString("GetUser") }
func (m *MockUserController) CreateUser(c *fiber.Ctx) error { return c.Status(201).SendString("CreateUser") }
func (m *MockUserController) UpdateUser(c *fiber.Ctx) error { return c.Status(200).SendString("UpdateUser") }
func (m *MockUserController) DeleteUser(c *fiber.Ctx) error { return c.Status(200).SendString("DeleteUser") }

func TestRegisterUserRoutes(t *testing.T) {
	app := fiber.New()
	mockCtl := &MockUserController{}

	RegisterUserRoutes(app, mockCtl)

	tests := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/users", 200, "GetUsers"},
		{"GET", "/users/1", 200, "GetUser"},
		{"POST", "/users", 201, "CreateUser"},
		{"PUT", "/users/1", 200, "UpdateUser"},
		{"DELETE", "/users/1", 200, "DeleteUser"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.url, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, tt.expectedCode, resp.StatusCode)

		body := make([]byte, resp.ContentLength)
		resp.Body.Read(body)
		resp.Body.Close()

		assert.Equal(t, tt.expectedBody, string(body))
	}
}
