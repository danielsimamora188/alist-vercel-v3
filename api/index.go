package handler

import (
	"net/http"
)

// Handler adalah fungsi utama untuk menangani request
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("AList is running"))
}
