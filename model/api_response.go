package model

// ApiResponse represents the structure of a generic API response.
type ApiResponse[T any] struct {
	Timestamp string `json:"timestamp"`
	QueryID   string `json:"queryId"`
	UserID    string `json:"userId"`
	SessionID string `json:"sessionId"`
	Data      T      `json:"data"`
}

//// ApiResponse represents the structure of a generic API response.
//type ApiResponse[T any] struct {
//	Metadata Metadata `json:"metadata"`
//	Data     T        `json:"data"` // T is a generic type
//}
//
//// Metadata represents the common metadata included in all responses.
//type Metadata struct {
//	Timestamp string `json:"timestamp"`
//	QueryID   string `json:"queryId"`
//	UserID    string `json:"userId"`
//	SessionID string `json:"sessionId"`
//}
