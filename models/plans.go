package models

type Plan struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}
