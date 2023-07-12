package routes

import (
	controllers "github.com/stripepaymentgateway/controllers"

	"github.com/gorilla/mux"
)


func PaymentRoutes() *mux.Router {
	router := mux.NewRouter();

	router.HandleFunc("/api/stripe/create-checkout-session", controllers.CreateCheckoutSession).Methods("POST");

	return router
}