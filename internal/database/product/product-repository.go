package product

import "github.com/Pahtll/hawy-go/internal/models"

type Repository interface {
	GetAll() ([]models.Product, error)
	GetById(id uint) (models.Product, error)
	GetByTitle(title string) (models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
	Delete(id uint) error
}
