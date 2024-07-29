package product

import "github.com/Pahtll/hawy-go/internal/models"

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetById(id uint64) (models.Product, error)
	GetByTitle(title string) (models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
	Delete(id uint64) error
}
