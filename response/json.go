package response

import (
	"encoding/json"
	"net/http"
)

// JSON is a helper function to write JSON responses
// to the client
// Example:
// response.JSON(w, http.StatusOK, todos)
func JSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		panic(err)
	}
}
