package database

import (
	"fmt"
	"github.com/Pahtll/hawy-go/internal/models"
	_ "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DB struct {
	*gorm.DB
}

var (
	DbHost     = os.Getenv("DB_HOST")
	DbPort     = os.Getenv("DB_PORT")
	DbUsername = os.Getenv("DB_USERNAME")
	DbDatabase = os.Getenv("DB_DATABASE")
	DbPassword = os.Getenv("DB_PASSWORD")
)

func New() (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s, port=%s",
		DbHost, DbUsername, DbPassword, DbDatabase, DbPort)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.CartItem{})
	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}
