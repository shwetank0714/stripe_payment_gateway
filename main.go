package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	routers "github.com/stripepaymentgateway/routes"
)

func main() {

	fmt.Println("Welcome to the Stripe Payment Gateway!")
	stripe.Key = "sk_test_51NSpYeSHczMHCorjG1Q5Fc9ZoYqOH4Fbw3xzU3XGMNP2fA7rTKrZIn6ZO21QZjtkRLxGD6FxmumNJuD899LXBPwS00BEyNcRnY"

	router := routers.PaymentRoutes()
	address := ":4242";
	log.Fatal(http.ListenAndServe(address, router))

}
