package postgres

import (
	"github.com/Pahtll/hawy-go/internal/database"
	"github.com/Pahtll/hawy-go/internal/database/user"
	"github.com/Pahtll/hawy-go/internal/models"
)

type userRepository struct {
	*database.DB
}

func New() (user.Repository, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	return &userRepository{
		db,
	}, nil
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	return users, result.Error
}

func (r *userRepository) GetById(id uint) (models.User, error) {
	var u models.User
	result := r.DB.First(&u, id)
	return u, result.Error
}

func (r *userRepository) GetByUsername(username string) (models.User, error) {
	var u models.User
	result := r.DB.Where("username = ?", username).First(&u)
	return u, result.Error
}

func (r *userRepository) Create(user models.User) error {
	result := r.DB.Create(&user)
	return result.Error
}

func (r *userRepository) Update(user models.User) error {
	result := r.DB.Save(&user)
	return result.Error
}

func (r *userRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	return result.Error
}
