package mappers

import (
	"comixhub/models"
	"comixhub/types"
)

func SubscriptionModelToSubscriptionType(subscriptionModel models.SubscriptionModel) types.Subscription {
	return types.Subscription{
		ID: subscriptionModel.SubscriptionID,
		User: types.User{
			ID:    subscriptionModel.UserID,
			Email: subscriptionModel.UserEmail,
			Name:  subscriptionModel.UserName,
		},
		Title: types.Title{
			ID:   subscriptionModel.TitleID,
			Name: subscriptionModel.TitleName,
		},
	}
}
