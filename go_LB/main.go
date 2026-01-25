package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Delayer
func main() {
	fmt.Println("DELAY: ", os.Getenv("DELAY"))
	delay, _ := time.ParseDuration(os.Getenv("DELAY"))
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9090"
	}

	fmt.Println("Delay: ", delay)
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(200)
		w.Write([]byte("ok\n"))
	})

	log.Printf("Backend on :%s with delay %s\n", port, delay)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Cmd to run: PORT=9001 DELAY=20ms go run main.go
