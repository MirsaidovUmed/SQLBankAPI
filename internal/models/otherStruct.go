package models

type ChangeBalanceRequest struct {
	Account Account `json:"account"`
	Amount  float64 `json:"amount"`
}

type TransferRequest struct {
	SenderPhone    string  `json:"sender"`
	RecipientPhone string  `json:"recipient"`
	Amount         float64 `json:"amount"`
}
