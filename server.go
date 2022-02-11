package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Go server starting on PORT:", PORT)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Route not found.", http.StatusNotFound)

		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)

		return
	}

	fmt.Fprintf(w, "hello-world")
}
