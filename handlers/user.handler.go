package handlers

import (
	"comixhub/types"
	"comixhub/utils"
	"encoding/json"
	"net/http"

	"comixhub/services"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	SubscriptionService *services.SubscriptionService
}

func NewUserHandler(subscriptionService services.SubscriptionService) *UserHandler {
	handler := UserHandler{
		SubscriptionService: &subscriptionService,
	}
	return &handler
}

func (handler *UserHandler) GetUserSubscriptionsHandler(w http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]

	// get array of subscriptions from service
	subscriptions, err := handler.SubscriptionService.GetUserSubscriptions(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RenderJSON(w, subscriptions)
}

func (handler *UserHandler) PostUserSubscriptionsHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	var subscriptionPost types.SubscriptionPostType
	if err := decoder.Decode(&subscriptionPost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subscription, err := handler.SubscriptionService.CreateUserSubscription(subscriptionPost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RenderJSON(w, subscription)
}
