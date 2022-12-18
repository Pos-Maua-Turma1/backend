package domain

import "time"

type User struct {
	id          string `json:"id" db:"id"`
	email 		string `json:"email" db:"user_id"`
	password 	string `json:"password" db:"password"`
	notes 		[]Note 
	createdAt   time.Time `json:"created_at" db:"created_at"`
	updatedAt   time.Time `json:"updated_at" db:"updated_at"`
}