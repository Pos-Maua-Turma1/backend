package dto

type LabelRequestDto struct {
	name          	string `json:"name" db:"id"`
	decription 		string `json:"decription" db:"user_id"`
}