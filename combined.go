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
			fmt.Fprintln(w, "Hello from the Go API!")
		})
		fmt.Println("Serving HTTP on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		fmt.Println("Hello, Everyday DevOps!")
	}
}
