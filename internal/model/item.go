package model

// Item represents the primary structure for items stored in the system.
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
