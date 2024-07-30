package mock

import (
	"github.com/Pahtll/hawy-go/internal/models"
	"testing"
)

func TestProductRepository(t *testing.T) {
	t.Parallel()
	var s = storage{
		products: []models.Product{
			{
				ID:    1,
				Title: "test",
				Price: 44.99,
			},
			{
				ID:    2,
				Title: "test2",
				Price: 1.99,
			},
			{
				ID:    3,
				Title: "test3",
				Price: 10.99,
			},
		},
	}
	t.Run("GetAll", func(t *testing.T) {
		products, err := s.GetAll()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("products: %v", products)
	})

	t.Run("GetById", func(t *testing.T) {
		product, err := s.GetById(1)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("product: %v", product)
	})

	t.Run("GetByTitle", func(t *testing.T) {
		products, err := s.GetByTitle("test")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("products: %v", products)
	})

	t.Run("Create", func(t *testing.T) {
		err := s.Create(models.Product{
			Title: "test",
			Price: 44.99,
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err := s.Update(models.Product{
			Title: "test",
			Price: 44.99,
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(1)
		if err != nil {
			t.Fatal(err)
		}
	})
}
