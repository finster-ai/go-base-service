package model

// ItemRequest represents the structure of an item as received in API requests.
type ItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
