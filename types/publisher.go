package types

type Publisher struct {
	ID       string     `json:"id"`
	Titles   *[]Title   `json:"titles"`
	Creators *[]Creator `json:"creators"`
}
