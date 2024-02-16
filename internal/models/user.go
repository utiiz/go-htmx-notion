package models

type User struct {
	ID        int    `json:"id_user" db:"id_user"`
	Lastname  string `json:"lastname" db:"lastname"`
	Firstname string `json:"firstname" db:"firstname"`
	Email     string `json:"email" db:"email"`
}
