package model

// ItemResponse represents the structure of an item sent in API responses.
type ItemResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
