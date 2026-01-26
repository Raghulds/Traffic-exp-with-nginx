package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Status string `json:"status"`
	Hash   string `json:"hash"`
}

func cpuWork() string {
	h := sha256.New()
	data := make([]byte, 32*1024) // 32KB
	rand.Read(data)
	h.Write(data)
	return string(h.Sum(nil))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		hash := cpuWork()

		// Random IO jitter
		time.Sleep(time.Duration(5+rand.Intn(15)) * time.Millisecond)

		// ---- slow path (1%) ----
		if rand.Intn(100) == 0 {
			time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
		}

		// error simulation
		if rand.Intn(500) == 0 {
			http.Error(w, "upstream failure", http.StatusInternalServerError)
			return
		}

		fmt.Println("Hash in Go: ", hash)
		resp := Response{
			Status: "ok",
			Hash:   hash,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		_ = start
	})

	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(5 * time.Millisecond)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("fast\n"))

		log.Printf("[FAST] took=%v\n", time.Since(start))
	})

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(500 * time.Millisecond)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("slow\n"))

		log.Printf("[SLOW] took=%v\n", time.Since(start))
	})

	log.Println("Starting go server on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
