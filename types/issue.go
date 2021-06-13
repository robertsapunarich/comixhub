package types

import (
	"time"
)

type Issue struct {
	ID          string        `json:"id"`
	Title       Title         `json:"title"`
	DiamondID   string        `json:"diamondId"`
	Price       float64       `json:"price"`
	ReleaseDate time.Time     `json:"releaseDate"`
	Creators    []CreatorRole `json:"creators"`
}
