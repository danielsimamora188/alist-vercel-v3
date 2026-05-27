package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. Cek apakah AList sudah menyala, jika belum, jalankan
	// Menggunakan /tmp karena itu satu-satunya folder yang bisa ditulis
	cmd := exec.Command("/tmp/alist", "server", "--no-prefix")
	cmd.Start()

	// 2. Beri waktu singkat agar biner siap
	time.Sleep(1 * time.Second)

	// 3. Proxy request ke port default AList
	target, _ := url.Parse("http://127.0.0.1:5244")
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
