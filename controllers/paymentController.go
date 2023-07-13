package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/stripe/stripe-go/v74/charge"
	"github.com/stripe/stripe-go/v74"

	helpers "github.com/stripepaymentgateway/helpers"
	models "github.com/stripepaymentgateway/models"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON payload from the request body
	var paymentInputDetail models.PaymentModel

	if err := json.NewDecoder(r.Body).Decode(&paymentInputDetail); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	// Set your Stripe secret key
	stripe.Key = "your_stripe_secret_key"

	// Create a card token param
	tokenParams := helpers.GetCardTokenParam(paymentInputDetail)

	// Generate the token
	cardToken, err := helpers.GenerateToken(tokenParams)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to tokenize card details")
		return
	}

	// Create a charge using the card token and payment details
	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(paymentInputDetail.PaymentAmount * 100),
		Currency:    stripe.String(paymentInputDetail.Currency),
		Description: stripe.String(paymentInputDetail.PaymentDescription),
		Source:      &stripe.PaymentSourceSourceParams{Token: stripe.String(cardToken.ID)},
	}

	// Perform the charge
	
	ch, err := charge.New(chargeParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Payment failed")
		return
	}

	// Payment was successful
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payment successful. Charge ID: %s", ch.ID)
}

// func CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
// 	domain := "http://localhost:4242/"

// 	params := &stripe.CheckoutSessionParams{
// 		LineItems: []*stripe.CheckoutSessionLineItemParams{
// 			{
// 				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
// 				Price:    stripe.String("{{PRICE_ID}}"),
// 				Quantity: stripe.Int64(1),
// 			},
// 		},
// 		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
// 		SuccessURL: stripe.String(domain + "/success.html"),
// 		CancelURL:  stripe.String(domain + "/cancel.html"),
// 	}

// 	s, err := session.New(params)

// 	if err != nil {
// 		log.Printf("session.New: %v", err)
// 	}

// 	http.Redirect(w, r, s.URL, http.StatusSeeOther)
// }
