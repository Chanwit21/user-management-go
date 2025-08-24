package model

import (
	"encoding/json"
	"testing"
)

func TestUserJSONMarshaling(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "John Doe",
		Username: "johndoe",
		Email:    "john@example.com",
		Phone:    "1234567890",
		Website:  "example.com",
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("failed to marshal user: %v", err)
	}

	expected := `{"id":1,"name":"John Doe","username":"johndoe","email":"john@example.com","phone":"1234567890","website":"example.com"}`
	if string(data) != expected {
		t.Errorf("expected JSON %s, got %s", expected, string(data))
	}
}

func TestUserJSONUnmarshaling(t *testing.T) {
	jsonStr := `{"id":2,"name":"Alice","username":"alice","email":"alice@example.com","phone":"9876543210","website":"alice.com"}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if user.ID != 2 || user.Name != "Alice" || user.Username != "alice" || user.Email != "alice@example.com" {
		t.Errorf("unmarshaled user does not match expected values: %+v", user)
	}
}
