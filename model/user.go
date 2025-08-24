package model

// User - simplified data model
type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Website  string `json:"website,omitempty"`
}
