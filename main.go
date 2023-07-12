package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/stripepaymentgateway/routes"
)

func main() {

	fmt.Println("Welcome to the Stripe Payment Gateway!")

	router := routers.PaymentRoutes()
	address := ":4242";
	log.Fatal(http.ListenAndServe(address, router))

}
