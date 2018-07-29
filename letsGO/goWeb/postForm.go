package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", handleRequest)
	http.ListenAndServe(":8000", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "   text/html;charset=utf8")
	if r.Method == "POST" {
		r.ParseForm()
		w.Write([]byte("username :" + r.FormValue("username") + "<br/>"))
		w.Write([]byte("<hr/>"))
		names := r.Form["username"]
		w.Write([]byte("that is two username: " + fmt.Sprintf("%v", names)))
		w.Write([]byte("<hr/>r.form's content : " + fmt.Sprintf("%v", r.Form)))
		w.Write([]byte("<hr/>r.PostForm's content : " + fmt.Sprintf("%v", r.Form)))
	} else {
		strBody := `<form action = ` + r.URL.RequestURI() + `"method = "post">
			username : <input name = "username" type = "text"/><br />
			userneme : <input name = "username" type = "text"/><br />
			<input id = "Submit1" type = "submit" value="submit" />
			</form>`
		w.Write([]byte(strBody))
		r.ParseForm()
	}
}
