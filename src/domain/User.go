package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	email 		*string `json:"email" db:"email"`
	password 	*string `json:"password" db:"password"`
	notes 		[]*Note `gorm:"foreignKey:userId"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err 
}