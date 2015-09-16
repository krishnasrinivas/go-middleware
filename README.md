# go-middleware
Go HTTP middleware

Example
-------


```
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnasrinivas/go-middleware"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("hello")
	next(w, r)
}

func HowHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("how")
	next(w, r)
}

func AreHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("are")
	next(w, r)
}

func YouHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("you")
	next(w, r)
}

func WtfHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gorilla!")
}

func main() {
	fmt.Println("starting..")
	r := mux.NewRouter()
	r.HandleFunc("/wtf", WtfHandler)
	mw := middleware.New(r, HelloHandler, HowHandler, AreHandler, YouHandler)
	s := &http.Server{
		Addr:    ":8000",
		Handler: mw,
	}

	s.ListenAndServe()
}


```
