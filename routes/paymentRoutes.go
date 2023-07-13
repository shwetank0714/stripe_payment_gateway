package routes

import (
	controllers "github.com/stripepaymentgateway/controllers"

	"github.com/gorilla/mux"
)


func PaymentRoutes() *mux.Router {
	router := mux.NewRouter();

	router.HandleFunc("/api/stripe/perform-payment", controllers.PaymentHandler).Methods("POST");

	return router;
}