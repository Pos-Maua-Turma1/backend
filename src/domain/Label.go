package domain

import "time"

type Label struct {
	id          	string `json:"id" db:"id"`
	name          	string `json:"name" db:"id"`
	decription 		string `json:"decription" db:"user_id"`
	notes			[]Note
	createdAt   	time.Time `json:"created_at" db:"created_at"`
	updatedAt   	time.Time `json:"updated_at" db:"updated_at"`
}