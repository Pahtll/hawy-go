package postgres

import (
	"fmt"
	"github.com/Pahtll/hawy-go/internal/database"
	"github.com/Pahtll/hawy-go/internal/database/product"
	"github.com/Pahtll/hawy-go/internal/models"
)

type repository struct {
	database.DB
}

func New(db database.DB) (product.Repository, error) {
	const op = "database.product.postgres.product-repository.New"

	if db.DB == nil {
		return nil, fmt.Errorf("%s: db is nil", op)
	}

	return &repository{
		db,
	}, nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	const op = "database.product.postgres.product-repository.GetAll"

	var products []models.Product

	result := r.DB.Find(&products)
	return products, fmt.Errorf("%s: %w", op, result.Error)
}

func (r *repository) GetById(id uint) (models.Product, error) {
	const op = "database.product.postgres.product-repository.GetById"

	var productEntity models.Product
	result := r.DB.First(&productEntity, id)
	if result.Error != nil {
		return models.Product{}, fmt.Errorf("%s: %w", op, result.Error)
	}
	return productEntity, nil
}

func (r *repository) GetByTitle(title string) (models.Product, error) {
	const op = "database.product.postgres.product-repository.GetByTitle"

	var productEntity models.Product
	result := r.DB.Where("title = ?", title).First(&productEntity)
	if result.Error != nil {
		return models.Product{}, fmt.Errorf("%s: %w", op, result.Error)
	}
	return productEntity, nil
}

func (r *repository) Create(product models.Product) error {
	const op = "database.product.postgres.product-repository.Create"

	result := r.DB.Create(&product)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", op, result.Error)
	}
	return nil
}

func (r *repository) Update(product models.Product) error {
	const op = "database.product.postgres.product-repository.Update"

	result := r.DB.Save(&product)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", op, result.Error)
	}
	return nil
}

func (r *repository) Delete(id uint) error {
	const op = "database.product.postgres.product-repository.Delete"

	result := r.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", op, result.Error)
	}
	return nil
}
