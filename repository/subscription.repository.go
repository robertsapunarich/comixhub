package repository

import (
	"comixhub/mappers"
	"comixhub/models"
	"comixhub/types"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	DatabaseConn *gorm.DB
}

func NewSubscriptionRepository(dbConn *gorm.DB) *SubscriptionRepository {
	repo := SubscriptionRepository{
		DatabaseConn: dbConn,
	}

	return &repo
}

func (repo *SubscriptionRepository) GetSubscriptionsByUserId(userId string) (*[]types.Subscription, error) {
	var subscriptionModels []models.SubscriptionModel
	var subscriptions []types.Subscription
	repo.DatabaseConn.Raw(
		`select 
			subscription.id subscription_id,
			user_id,
			"user".email user_email,
			"user"."name" user_name,
			title_id,
			title."name" title_name
		from
			user_title_subscription subscription 
		join "user"
			on "user".id = subscription.user_id
		join title 
			on title.id = subscription.title_id
		where user_id = ?`,
		userId,
	).Scan(&subscriptionModels)

	for _, subscriptionModel := range subscriptionModels {
		subscriptions = append(subscriptions, mappers.SubscriptionModelToSubscriptionType(subscriptionModel))
	}

	return &subscriptions, nil
}
