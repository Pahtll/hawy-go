package mock

import (
	"fmt"
	"github.com/Pahtll/hawy-go/internal/database/user"
	"github.com/Pahtll/hawy-go/internal/models"
)

type Storage struct {
	users []models.User
}

func New() user.Repository {
	return &Storage{
		users: []models.User{},
	}
}

func (m *Storage) GetAll() ([]models.User, error) {
	const op = "mock.user.GetAll"
	if m.users == nil {
		return []models.User{}, fmt.Errorf("%s: Users not found", op)
	}
	return m.users, nil
}

func (m *Storage) GetById(id uint64) (models.User, error) {
	const op = "mock.user.GetById"
	if &m.users[id] == nil {
		return models.User{}, fmt.Errorf("%s: User not found", op)
	}
	return m.users[id], nil
}

func (m *Storage) GetByUsername(username string) (models.User, error) {
	const op = "mock.user.GetByUsername"
	for _, u := range m.users {
		if u.Username == username {
			return u, nil
		}
	}
	return models.User{}, fmt.Errorf("%s: User not found", op)
}

func (m *Storage) Create(user models.User) error {
	const op = "mock.user.Create"
	m.users = append(m.users, user)
	return nil
}

func (m *Storage) Update(user models.User) error {
	const op = "mock.user.Update"
	m.users[user.ID] = user
	return nil
}

func (m *Storage) Delete(id uint64) error {
	const op = "mock.user.Delete"
	m.users = append(m.users[:id], m.users[id+1:]...)
	return nil
}
