package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Arahkan ke biner alist yang berjalan di local
	target, _ := url.Parse("http://127.0.0.1:5244")
	proxy := httputil.NewSingleHostReverseProxy(target)
	
	// Jalankan proxy
	proxy.ServeHTTP(w, r)
}
