package models

type Project struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Icon string `json:"icon" db:"icon"`
	Url  string `json:"url" db:"url"`
}
