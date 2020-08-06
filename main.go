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

	mux := http.NewServeMux()

	v1Mux := http.NewServeMux()
	v1Mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "v1 Profile")
	})
	v1Mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "v1 Posts")
	})

	v2Mux := http.NewServeMux()
	v2Mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "v2 Profile")
	})
	v2Mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "v2 Posts")
	})

	mux.Handle("/v1/", http.StripPrefix("/v1", v1Mux))
	mux.Handle("/v2/", http.StripPrefix("/v2", v2Mux))
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/profile", getProfile)

	// http.Handle("/", http.FileServer(http.Dir("/tmp")))

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

	if err := http.ListenAndServe(":8080", mux); err != nil {
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