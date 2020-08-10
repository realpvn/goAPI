// Good Practices
// Struct name should start with lowercase
// Json tags should start with uppercase

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// creating a struct to store each value to be returned
type helloWorld struct {
	Message string `json:"message"`
}

// creating a variable of type helloWorld
type hello []helloWorld

func showMessage(w http.ResponseWriter, r *http.Request) {
	// hard coding values to the json
	// which in most cases will be fetched by database
	admins := hello{
		helloWorld{Message: "Hello World"},
	}

	fmt.Println("All Admins")
	json.NewEncoder(w).Encode(admins)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func handleRequest() {
	//handling api requests
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hi", showMessage)

	//listening on localhost port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
