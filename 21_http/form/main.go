package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	if method := req.Method; method == http.MethodGet {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Fprintf(w, "err %v", err)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "err %v", err)
		}
	} else {
		println("req.Form:", req.Form)
		println("req.Form:", req.Form == nil)
		println("req.PostForm:", req.PostForm)
		println("FormValue:", req.FormValue("username"))
	}
}
