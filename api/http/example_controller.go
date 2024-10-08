package http

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	model2 "go-base-service/internal/model"
	"go-base-service/internal/service"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var (
	itemService *service.ItemService
	once        sync.Once // To ensure the service is only initialized once
)

func InitController() {
	once.Do(func() {
		itemService = service.NewItemService() // Initialize ItemService here
	})
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	logrus.Info("/items API Endpoint Reached")
	vars := mux.Vars(r)
	id := vars["id"]

	item, exists := itemService.FindItem(id)
	if !exists {
		logrus.Error("ExampleCallReturnsEmpty GRPC Endpoint Reached")
		generateErrorResponse(w, http.StatusNotFound, "Item not found", 404)
		return
	}

	responseData := model2.ItemResponse{
		ID:   item.ID,
		Name: item.Name,
	}

	generateSuccessResponse(w, responseData, "unique-query-id-get")
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var itemRequest model2.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&itemRequest); err != nil {
		generateErrorResponse(w, http.StatusBadRequest, err.Error(), 400)
		return
	}

	item := model2.Item{
		ID:   itemRequest.ID, // Placeholder: Replace with actual ID generation logic
		Name: itemRequest.Name,
	}

	createdItem := itemService.AddItem(item)

	responseData := model2.ItemResponse{
		ID:   createdItem.ID,
		Name: createdItem.Name,
	}

	generateSuccessResponse(w, responseData, "unique-query-id-create", http.StatusCreated)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !itemService.RemoveItem(id) {
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

	response := model2.ApiResponse[interface{}]{
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
	errorResponse := model2.ApiErrorResponse{
		Errors: []model2.ErrorDetail{
			{
				Code:    code,
				Message: message,
			},
		},
	}
	json.NewEncoder(w).Encode(errorResponse)
}
