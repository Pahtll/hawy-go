package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"unique;not null;type:varchar(100);column:title"`
	Price float64
}
