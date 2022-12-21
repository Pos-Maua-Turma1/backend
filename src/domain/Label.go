package domain

import (

	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	name          	string 	 `json:"name" db:"id"`
	decription 		string 	 `json:"decription" db:"decription"`
	notes			[]*Note  `gorm:"many2many:label_note"`
}

func MigrateLabel(db *gorm.DB) error {
	err := db.AutoMigrate(&Label{})
	return err 
}