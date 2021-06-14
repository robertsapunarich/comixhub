package services

import (
	"comixhub/repository"
	"comixhub/types"
	"log"
)

type SubscriptionService struct {
	SubscriptionRepo *repository.SubscriptionRepository
}

func (subService *SubscriptionService) GetUserSubscriptions(userId string) ([]types.Subscription, error) {
	log.Print("getting subscription by user id")
	// Get subscriptions by user id

	// Handle error
	// Convert result set into subscriptions with a mapper
	// Handle error
	// Return array of subscriptions
	return nil, nil
}
