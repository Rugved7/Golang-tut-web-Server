package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Modules Section in Golang")
	// call the function
	greeter()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	// Listening on port
	log.Fatal(http.ListenAndServe(":8080", router))
}

func greeter() {
	fmt.Println("Hello from the Golang")
}

// Some new and Serious Shit
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang</h1>"))

}
