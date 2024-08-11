package mock

import (
	"fmt"
	"github.com/Pahtll/hawy-go/internal/database/product"
	"github.com/Pahtll/hawy-go/internal/models"
)

type storage struct {
	products []models.Product
}

func New() product.Repository {
	return &storage{
		products: []models.Product{},
	}
}

func (s *storage) GetAll() ([]models.Product, error) {
	const op = "database.product.mock.product-repository.GetAll"

	if s.products == nil {
		return []models.Product{}, fmt.Errorf("%s:Products not found", op)
	}

	return s.products, nil
}

func (s *storage) GetById(id uint) (models.Product, error) {
	const op = "database.product.mock.product-repository.GetById"

	if s.products == nil {
		return models.Product{}, fmt.Errorf("%s:Products not found", op)
	}

	if &s.products[id-1] == nil {
		return models.Product{}, fmt.Errorf("%s:Product not found", op)
	}

	return s.products[id-1], nil
}

func (s *storage) GetByTitle(title string) (models.Product, error) {
	const op = "database.product.mock.product-repository.GetByTitle"

	if s.products == nil {
		return models.Product{}, fmt.Errorf("%s:Products not found", op)
	}

	for _, p := range s.products {
		if p.Title == title {
			return p, nil
		}
	}

	return models.Product{}, fmt.Errorf("%s:Product not found", op)
}

func (s *storage) Create(product models.Product) error {
	const op = "database.product.mock.product-repository.Create"

	if s.products == nil {
		s.products = []models.Product{}
	}

	for _, p := range s.products {
		if p == product {
			return fmt.Errorf("%s:Product already exists", op)
		}
	}

	s.products = append(s.products, product)
	return nil
}

func (s *storage) Update(product models.Product) error {
	const op = "database.product.mock.product-repository.Update"

	if s.products == nil {
		return fmt.Errorf("%s:Products not found", op)
	}

	for i, p := range s.products {
		if p.ID == product.ID {
			s.products[i] = product
			return nil
		}
	}

	return fmt.Errorf("%s:Product not found", op)
}

func (s *storage) Delete(id uint) error {
	const op = "database.product.mock.product-repository.Delete"

	if s.products == nil {
		return fmt.Errorf("%s:Products not found", op)
	}

	for i, p := range s.products {
		if p.ID == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("%s:Product not found", op)
}
