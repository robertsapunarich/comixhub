package types

type User struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Email         string          `json:"email"`
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
}
