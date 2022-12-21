package domain

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	userId 		string
	noteBody 	string `json:"note_body" db:"note_body"`
	labels 		[]*Label `gorm:"many2many:label_note"`
}

func MigrateNote(db *gorm.DB) error {
	err := db.AutoMigrate(&Note{})
	return err 
}