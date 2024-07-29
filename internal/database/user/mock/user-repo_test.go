package mock

import (
	"github.com/Pahtll/hawy-go/internal/models"
	"testing"
)

func TestStorage_GetAllFailed(t *testing.T) {
	t.Parallel()
	var s Storage

	t.Run("GetAll", func(t *testing.T) {
		users, err := s.GetAll()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("users: %v", users)
	})

	t.Run("GetAll", func(t *testing.T) {
		s.users = []models.User{
			{
				ID:           1,
				Username:     "testuser",
				Email:        "testemail",
				PasswordHash: "123a-4123-asdv-12sd",
				Role:         0,
				Cart:         []models.CartItem{},
			},
			{
				ID:           2,
				Username:     "usr2",
				Email:        "mailsd",
				PasswordHash: "123d-4123-asdv-12sd",
				Role:         2,
				Cart:         []models.CartItem{},
			},
		}
		users, err := s.GetAll()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("users: %v", users)
	})
}

func TestStorage_GetById(t *testing.T) {
	t.Parallel()
	var s Storage

	t.Run("GetById", func(t *testing.T) {
		s.users = []models.User{
			{
				ID:           1,
				Username:     "testuser",
				Email:        "testemail",
				PasswordHash: "123a-4123-asdv-12sd",
				Role:         0,
				Cart:         []models.CartItem{},
			},
			{
				ID:           2,
				Username:     "usr2",
				Email:        "mailsd",
				PasswordHash: "123d-4123-asdv-12sd",
				Role:         2,
				Cart:         []models.CartItem{},
			},
		}
		user, err := s.GetById(1)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("user: %v", user)
	})
}

func TestStorage_GetByUsername(t *testing.T) {
	t.Parallel()
	var s Storage
	t.Run("GetByUsername", func(t *testing.T) {
		s.users = []models.User{
			{
				ID:           1,
				Username:     "testuser",
				Email:        "testemail",
				PasswordHash: "123a-4123-asdv-12sd",
				Role:         0,
				Cart:         []models.CartItem{},
			},
			{
				ID:           2,
				Username:     "usr2",
				Email:        "mailsd",
				PasswordHash: "123d-4123-asdv-12sd",
				Role:         2,
				Cart:         []models.CartItem{},
			},
		}
		user, err := s.GetByUsername("usr2")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("user: %v", user)
	})
}

func TestStorage_Create(t *testing.T) {
	t.Parallel()
	var s Storage
	t.Run("Create", func(t *testing.T) {
		err := s.Create(models.User{
			Username:     "test",
			Email:        "testmail",
			PasswordHash: "123a-4123-asdv-12sd",
			Role:         0,
			Cart:         []models.CartItem{},
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestStorage_Update(t *testing.T) {
	t.Parallel()
	var s Storage
	t.Run("Update", func(t *testing.T) {

		s.users = []models.User{
			{
				ID:           0,
				Username:     "test",
				Email:        "testmail",
				PasswordHash: "123a-4123-asdv-12sd",
				Role:         0,
				Cart:         []models.CartItem{},
			},
		}
		err := s.Update(models.User{
			ID:           0,
			Username:     "changed",
			Email:        "testmail",
			PasswordHash: "123a-4123-asdv-12sd",
			Role:         0,
			Cart:         []models.CartItem{},
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestStorage_DeleteFail(t *testing.T) {
	t.Parallel()
	var s Storage
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(1)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestStorage_Delete(t *testing.T) {
	t.Parallel()
	var s Storage
	t.Run("Delete", func(t *testing.T) {

		s.users = []models.User{
			{
				ID:           0,
				Username:     "testuser",
				Email:        "testemail",
				PasswordHash: "123a-4123-asdv-12sd",
				Role:         0,
				Cart:         []models.CartItem{},
			},
			{
				ID:           1,
				Username:     "usr2",
				Email:        "mailsd",
				PasswordHash: "123d-4123-asdv-12sd",
				Role:         2,
				Cart:         []models.CartItem{},
			},
		}
		err := s.Delete(0)
		if err != nil {
			t.Fatal(err)
		}
	})
}
