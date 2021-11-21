package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello med"))
}

func showPatient(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	i, err := fmt.Fprintf(w, "Show patient number %d", id)
	log.Println(i, err)
}

func createPatient(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Add("name", "med")

		w.Header().Del("name")
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// this is equivalent to
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("createPatient"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/patient", showPatient)
	mux.HandleFunc("/patient/create", createPatient)

	log.Println("Starting server ....")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
