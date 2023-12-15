package models

type Transfer struct {
	Sender    *Account
	Recipient *Account
	Amount    float64
}
