package models

type Account struct {
	Id          int     `json:"id"`
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Address     string  `json:"address"`
	Balance     float64 `json:"balance"`
}
