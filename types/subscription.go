package types

type Subscription struct {
	ID    string `json:"id"`
	User  User   `json:"user"`
	Title Title  `json:"title"`
}
