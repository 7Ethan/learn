package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func echo(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Echo error!"))
		return
	}

	writeLen, err := w.Write([]byte(msg))
	if err != nil || writeLen != len(msg) {
		log.Println("That's some error while write httpbody..")
	}
}

func TimeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		time := time.Since(start)
		log.Println("Cost time is : ", time)
	})
}

func main() {
	http.Handle("/", TimeMiddleWare(http.HandlerFunc(echo)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("server error!")
	}
}
