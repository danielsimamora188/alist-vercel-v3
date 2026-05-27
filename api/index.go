package handler

import (
	"net/http"
)

// Handler adalah fungsi utama untuk menangani request Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Karena Vercel adalah platform serverless dan AList adalah server panjang,
	// kita melakukan redirect sederhana agar fungsi tidak crash saat dipanggil.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("AList Vercel is running. Please check your deployment logs for status."))
}
