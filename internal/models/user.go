package models

import "gorm.io/gorm"

type UserRole uint8

const (
	Default UserRole = iota
	Moderator
	Admin
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"unique;not null;type:varchar(100);column:username"`
	PasswordHash string
	Email        string `gorm:"column:email;type:varchar(100);not null;unique"`
	Role         UserRole
	Cart         []CartItem
}
