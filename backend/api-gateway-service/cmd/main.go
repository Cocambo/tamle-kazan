package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API Gateway is running!")
	})

	fmt.Println("Listening on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
