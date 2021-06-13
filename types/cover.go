package types

type Cover struct {
	ID          string        `json:"id"`
	IssueID     string        `json:"issueId"`
	Creators    []CreatorRole `json:"creators"`
	Description string        `json:"description"`
}
