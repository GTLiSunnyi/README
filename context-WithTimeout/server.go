package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler start...")

	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("responsed to client.")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("timeout, canceled")
	}
}
