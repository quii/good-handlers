package main

import (
	"fmt"
	"github.com/gorilla/mux"
	goodhandlers "github.com/quii/good-handlers"
	"github.com/quii/good-handlers/kyc"
	"log"
	"net/http"
)

func main() {
	createKYCHandler := goodhandlers.New(
		kyc.CreateKYC,
		kyc.DecodeRequest,
		kyc.EncodeResponse,
	)

	router := mux.NewRouter()

	// goodhandlers just uses the go standard lib interface, so you can plug it into gorilla/chi/stdlib routers like normal
	router.Handle("/kyc", createKYCHandler)
	router.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
