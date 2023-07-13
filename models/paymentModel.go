package models

type PaymentModel struct {
	CardNumber         string `json:"card_number"`
	ExpiryMonth        string `json:"expiry_month"`
	ExpiryYear         string `json:"expiry_year"`
	CVV                string `json:"cvv"`
	PaymentAmount      int64 `json:"payment_amount"`
	Currency           string `json:"currency"`
	PaymentDescription string `json:"payment_desctiption"`
}
