package handler
import (
    "net/http"
    "os/exec"
)
func Handler(w http.ResponseWriter, r *http.Request) {
    // Kode ini hanya memicu Vercel untuk menjalankan binary alist 
    // yang akan kita unduh otomatis saat build
    http.Redirect(w, r, "https://github.com/alist-org/alist", http.StatusMovedPermanently)
}
