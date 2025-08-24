package service

import (
	"testing"
	"user-management-go/model"

	"github.com/stretchr/testify/assert"
)

func TestUserService_CRUD(t *testing.T) {
	// Initialize service without hitting the real HTTP API
	s := &UserService{
		users:  make(map[uint64]model.User),
		nextID: 1,
	}

	// --- CREATE ---
	user := model.User{Name: "John Doe", Email: "john@example.com"}
	created := s.Create(user)
	assert.Equal(t, uint64(1), created.ID)
	assert.Equal(t, "John Doe", created.Name)

	// --- GET ALL ---
	all := s.GetAll()
	assert.Len(t, all, 1)

	// --- GET BY ID ---
	got, ok := s.GetByID(created.ID)
	assert.True(t, ok)
	assert.Equal(t, "John Doe", got.Name)

	// --- UPDATE ---
	updatedUser := model.User{Name: "Jane Doe", Email: "jane@example.com"}
	updated, err := s.Update(created.ID, updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", updated.Name)

	// --- UPDATE NON-EXISTENT ---
	_, err = s.Update(999, updatedUser)
	assert.Error(t, err)

	// --- DELETE ---
	deleted := s.Delete(created.ID)
	assert.True(t, deleted)

	// --- DELETE NON-EXISTENT ---
	deleted = s.Delete(999)
	assert.False(t, deleted)

	// --- GET ALL AFTER DELETE ---
	all = s.GetAll()
	assert.Len(t, all, 0)
}
