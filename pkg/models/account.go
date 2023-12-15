package models

type Account struct {
	Name        string
	Balance     float64
	PhoneNumber string
	City        *City
}
