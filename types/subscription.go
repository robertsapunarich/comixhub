package types

type Subscription struct {
	ID    string `json:"id"`
	User  User   `json:"user"`
	Title Title  `json:"title"`
}

type SubscriptionPostType struct {
	UserID  string `json:"userId"`
	TitleID string `json:"titleId"`
}
