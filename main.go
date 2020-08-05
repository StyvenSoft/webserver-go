package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/welcome", welcome)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} 	
}

// func ListenAndServe(addr string, handler Handler) error

func welcome(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome to the GO server..")
}