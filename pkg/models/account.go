package models

type Account struct {
	Id          int
	Name        string
	Balance     float64
	Password    string
	PhoneNumber string
	City        *City
}
