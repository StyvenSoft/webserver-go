package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/welcome", welcome)
}

func welcome(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome to th GO server..")
}