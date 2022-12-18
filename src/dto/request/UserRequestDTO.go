package dto

type UserRequestDto struct {
	email 		string `json:"email" db:"user_id"`
	password 	string `json:"password" db:"note_body"`
}