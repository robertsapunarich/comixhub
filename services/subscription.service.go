package services

import (
	"comixhub/repository"
	"comixhub/types"
	"log"
)

type SubscriptionService struct {
	SubscriptionRepo *repository.SubscriptionRepository
}

func (subService *SubscriptionService) GetUserSubscriptions(userId string) (*[]types.Subscription, error) {
	log.Print("getting subscription by user id")
	// Get subscriptions by user id
	subscriptions, err := subService.SubscriptionRepo.GetSubscriptionsByUserId(userId)
	// Handle error
	if err != nil {
		return nil, err
	}

	// Convert result set into subscriptions with a mapper
	// Return array of subscriptions
	return subscriptions, nil
}

func (subservice *SubscriptionService) CreateUserSubscription(subscriptionPost types.SubscriptionPostType) (*[]types.Subscription, error) {
	log.Print("creating subscription...")

	return nil, nil
}
