package domain

import "time"

type Note struct {
	id          string `json:"id" db:"id"`
	userId 		string `json:"user_id" db:"user_id"`
	noteBody 	string `json:"note_body" db:"note_body"`
	labels 		[]Label
	createdAt   time.Time `json:"created_at" db:"created_at"`
	updatedAt   time.Time `json:"updated_at" db:"updated_at"`
}