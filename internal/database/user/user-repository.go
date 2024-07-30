package user

import "github.com/Pahtll/hawy-go/internal/models"

type Repository interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (models.User, error)
	GetByUsername(username string) (models.User, error)
	Create(user models.User) error
	Update(user models.User) error
	Delete(id uint) error
}
