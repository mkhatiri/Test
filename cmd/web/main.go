package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/patient", showPatient)
	mux.HandleFunc("/patient/create", createPatient)

	log.Println("Starting server ....")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
