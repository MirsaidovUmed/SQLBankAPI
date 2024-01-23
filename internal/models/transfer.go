package models

type Transfer struct {
	Id        int
	Sender    *Account
	Recipient *Account
	Amount    float64
}
