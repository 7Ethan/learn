package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	http.Handle("/", useLoggingHandler(handle()))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}
}

func handle() http.Handler {
	return http.HandlerFunc(myHandle)
}

func myHandle(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello gorilla"))
}

func useLoggingHandler(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
