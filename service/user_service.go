package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"user-management-go/model"
)

type UserService struct {
	sync.RWMutex
	users  map[uint64]model.User
	nextID uint64
}

func NewUserService() *UserService {
	s := &UserService{
		users:  make(map[uint64]model.User),
		nextID: 1,
	}

	// Fetch initial users from JSONPlaceholder
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err == nil {
		defer resp.Body.Close()
		var users []model.User
		if err := json.NewDecoder(resp.Body).Decode(&users); err == nil {
			for _, u := range users {
				s.Create(u)
			}
		}
	}

	return s
}

func (s *UserService) GetAll() []model.User {
	s.RLock()
	defer s.RUnlock()
	result := make([]model.User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result
}

func (s *UserService) GetByID(id uint64) (model.User, bool) {
	s.RLock()
	defer s.RUnlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *UserService) Create(u model.User) model.User {
	s.Lock()
	defer s.Unlock()
	u.ID = s.nextID
	s.nextID++
	s.users[u.ID] = u
	return u
}

func (s *UserService) Update(id uint64, u model.User) (model.User, error) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.users[id]
	if !ok {
		return model.User{}, errors.New("not found")
	}
	u.ID = id
	s.users[id] = u
	return u, nil
}

func (s *UserService) Delete(id uint64) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}
