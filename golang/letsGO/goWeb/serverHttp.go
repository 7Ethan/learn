package main

import (
	"net/http"
)

func main() {
	hadler := &HttpHandler{}
	http.ListenAndServe(":8000", hadler)
}

type HttpHandler struct {
}

func (this *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>在 ServeHTTP 里</h1>"))
	w.Write([]byte(r.URL.Path))
}
