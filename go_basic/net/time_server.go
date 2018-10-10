package main

import (
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (t *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time := time.Now().Format(t.format)
	w.Write([]byte("Now time is : " + time))
}

func redirect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path[len(r.URL.Path)-1] == '/' {
			http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)
	rd := redirect(mux)
	http.ListenAndServe(":3000", rd)
}
