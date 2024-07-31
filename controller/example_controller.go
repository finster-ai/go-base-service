package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go-base-service/model"
	"go-base-service/service"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	item, exists := service.FindItem(id)
	if !exists {
		generateErrorResponse(w, http.StatusNotFound, "Item not found", 404)
		return
	}

	responseData := model.ItemResponse{
		ID:   item.ID,
		Name: item.Name,
	}

	generateSuccessResponse(w, responseData, "unique-query-id-get")
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var itemRequest model.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&itemRequest); err != nil {
		generateErrorResponse(w, http.StatusBadRequest, err.Error(), 400)
		return
	}

	item := model.Item{
		ID:   itemRequest.ID, // Placeholder: Replace with actual ID generation logic
		Name: itemRequest.Name,
	}

	createdItem := service.AddItem(item)

	responseData := model.ItemResponse{
		ID:   createdItem.ID,
		Name: createdItem.Name,
	}

	generateSuccessResponse(w, responseData, "unique-query-id-create", http.StatusCreated)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !service.RemoveItem(id) {
		generateErrorResponse(w, http.StatusNotFound, "Item not found", 404)
		return
	}

	generateSuccessResponse(w, nil, "unique-query-id-delete", http.StatusNoContent)
}

// generateSuccessResponse writes a success response with metadata and data
func generateSuccessResponse(w http.ResponseWriter, data interface{}, queryID string, status ...int) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	w.WriteHeader(statusCode)

	response := model.ApiResponse[interface{}]{
		Timestamp: time.Now().Format(time.RFC3339),
		QueryID:   queryID,
		UserID:    "user@example.com", // Placeholder: Replace with actual user ID logic
		SessionID: "session-id",       // Placeholder: Replace with actual session ID logic
		Data:      data,
	}

	json.NewEncoder(w).Encode(response)
}

// generateErrorResponse writes an error response with a standardized format
func generateErrorResponse(w http.ResponseWriter, statusCode int, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := model.ApiErrorResponse{
		Errors: []model.ErrorDetail{
			{
				Code:    code,
				Message: message,
			},
		},
	}
	json.NewEncoder(w).Encode(errorResponse)
}
