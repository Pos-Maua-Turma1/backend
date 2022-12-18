package dto

type NoteRequestDto struct {
	userId 		string `json:"user_id" db:"user_id"`
	noteBody 	string `json:"note_body" db:"note_body"`
}