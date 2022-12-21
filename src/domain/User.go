package domain

import (
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	id       int    `json:"id"`
	email    string `json:"email"`
	password string `json:"password"`
	// notes    []*Note `gorm:"foreignKey:userId"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
