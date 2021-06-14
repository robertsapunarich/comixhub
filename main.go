package main

import (
	"comixhub/handlers"
	"comixhub/repository"
	"comixhub/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	dbConnString := "host=localhost user=robert_sapunarich dbname=comixhub port=5432 sslmode=disable"
	subscriptionRepo := repository.NewSubscriptionRepository(dbConnString)
	subscriptionService := services.SubscriptionService{
		SubscriptionRepo: subscriptionRepo,
	}
	router := mux.NewRouter()
	userHandler := handlers.NewUserHandler(subscriptionService)

	router.HandleFunc("/user/{id}/subscriptions", userHandler.GetUserSubscriptionsHandler).Methods("GET")
	router.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), router))
}
