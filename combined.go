// combined.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	api := flag.Bool("api", false, "Run as HTTP API")
	flag.Parse()

	if *api {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintln(w, "Hello from the Go API!")
			if err != nil {
				// Handle the error, e.g., log it or exit
				log.Fatal("Error writing response")
			}
		})
		fmt.Println("Serving HTTP on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		fmt.Println("Hello, Everyday DevOps!")
	}
}
