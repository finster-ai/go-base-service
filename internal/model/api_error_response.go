package model

// ErrorDetail represents the structure of an individual error.
type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ApiErrorResponse represents the structure of the error response.
type ApiErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}
