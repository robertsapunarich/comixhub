package main

import (
	"comixhub/handlers"
	"comixhub/repository"
	"comixhub/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDbConn() *gorm.DB {
	dbConnString := "host=localhost user=robertsapunarich dbname=comixhub port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to initialize db connection: %s", dbConnString))
	}

	return db
}

func main() {
	dbConn := setupDbConn()
	subscriptionRepo := repository.NewSubscriptionRepository(dbConn)
	subscriptionService := services.SubscriptionService{
		SubscriptionRepo: subscriptionRepo,
	}
	router := mux.NewRouter()
	userHandler := handlers.NewUserHandler(subscriptionService)

	router.HandleFunc("/user/{id}/subscriptions", userHandler.GetUserSubscriptionsHandler).Methods("GET")
	router.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), router))
}
