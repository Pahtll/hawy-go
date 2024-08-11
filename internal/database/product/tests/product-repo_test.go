package tests

import (
	"github.com/Pahtll/hawy-go/internal/database/product/mock"
	"github.com/Pahtll/hawy-go/internal/models"
	"testing"
)

func TestProductRepository(t *testing.T) {
	var s = mock.New()

	err := s.Create(models.Product{
		ID:    1,
		Title: "test2",
		Price: 99.99,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Create", func(tt *testing.T) {
		err := s.Create(models.Product{
			ID:    2,
			Title: "test1",
			Price: 100.99,
		})
		if err != nil {
			tt.Fatal(err)
		}
	})

	t.Run("GetAll", func(tt *testing.T) {
		tt.Parallel()

		products, err := s.GetAll()
		if err != nil {
			tt.Fatal(err)
		}
		tt.Logf("products: %v", products)
	})

	t.Run("GetById", func(tt *testing.T) {
		tt.Parallel()

		product, err := s.GetById(1)
		if err != nil {
			tt.Fatal(err)
		}
		tt.Logf("product: %v", product)
	})

	t.Run("GetByTitle", func(tt *testing.T) {
		tt.Parallel()

		products, err := s.GetByTitle("test2")
		if err != nil {
			tt.Fatal(err)
		}
		tt.Logf("products: %v", products)
	})

	t.Run("Update", func(tt *testing.T) {
		tt.Parallel()

		err := s.Update(models.Product{
			ID:    1,
			Title: "test",
			Price: 44.99,
		})
		if err != nil {
			tt.Fatal(err)
		}
	})

	t.Run("Delete", func(tt *testing.T) {
		err := s.Create(models.Product{
			ID:    3,
			Title: "test12",
			Price: 1.99,
		})

		err = s.Delete(3)
		if err != nil {
			tt.Fatal(err)
		}
	})
}
