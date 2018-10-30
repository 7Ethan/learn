package main

import (
	"expvar"
	"fmt"
	"net/http"
	"time"
)

var (
	visits              = expvar.NewInt("visits")
	stasts              = expvar.NewMap("http")
	req_succ, req_faild expvar.Int
	start               = time.Now()
)

func init() {
	stasts.Set("request_success", &req_succ)
	stasts.Set("request_faild", &req_faild)
}

func calculateUptime() interface{} {
	return time.Since(start).String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	req_succ.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func hello(w http.ResponseWriter, r *http.Request) {
	req_faild.Add(1)
	w.Write([]byte("test faild"))
}

func main() {
	expvar.Publish("uptime", expvar.Func(calculateUptime))
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
