package models

type Account struct {
	Id          int
	Name        string
	Balance     float64
	PhoneNumber string
	City        *City
}
