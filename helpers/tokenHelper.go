package helpers

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/token"

	models "github.com/stripepaymentgateway/models"
)

func GetCardTokenParam(paymentInputDetail models.PaymentModel) *stripe.TokenParams{
	return &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(paymentInputDetail.CardNumber),
			ExpMonth: stripe.String(paymentInputDetail.ExpiryMonth),
			ExpYear:  stripe.String(paymentInputDetail.ExpiryYear),
			CVC:      stripe.String(paymentInputDetail.CVV),
		},
	}
}

func GenerateToken(tokenParams *stripe.TokenParams) (*stripe.Token, error) {
	return token.New(tokenParams)
}