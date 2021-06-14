package handlers

import (
	"comixhub/utils"
	"net/http"

	"comixhub/services"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	SubscriptionService *services.SubscriptionService
}

func NewUserHandler(subscriptionService services.SubscriptionService) UserHandler {
	handler := UserHandler{
		SubscriptionService: &subscriptionService,
	}
	return handler
}

func (handler *UserHandler) GetUserSubscriptionsHandler(w http.ResponseWriter, req *http.Request) {
	subscriptionService := services.SubscriptionService{}
	userId := mux.Vars(req)["id"]

	// get array of subscriptions from service
	subscriptions, err := subscriptionService.GetUserSubscriptions(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RenderJSON(w, subscriptions)
}

func PostUserSubscriptionHandler(w http.ResponseWriter, req *http.Request) {

}
