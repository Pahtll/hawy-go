package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID    uint64
	ProductID uint64
	Product   Product `gorm:"foreignKey:ProductID"`
}
