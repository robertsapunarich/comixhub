package types

type Title struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Publisher Publisher  `json:"publisher"`
	Creators  *[]Creator `json:"creators"`
}
