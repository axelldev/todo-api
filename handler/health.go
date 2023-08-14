package handler

import (
	"net/http"

	"github.com/axelldev/todo-api/response"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "OK")
}
