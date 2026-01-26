package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/copy", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Slow upstream (Any delay after NGINX accepting the request)
		n, _ := io.Copy(io.Discard, r.Body)
		fmt.Printf("Copied bytes: %d\n", n)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Copy body success!"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
