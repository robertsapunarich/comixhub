package types

type Creator struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatorRole struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
