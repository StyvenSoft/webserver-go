package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type profile struct {
	Name string
	Hobbies []string
}

func main()  {
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/profile", getProfile)

	// http.Handle("/", http.FileServer(http.Dir("/tmp")))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} 	
}

func getProfile(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	profile := profile{
		Name: "Steveen Echeverri",
		Hobbies: []string{"Programming", "Design"},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func ListenAndServe(addr string, handler Handler) error

func welcome(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome to the GO server..")
}