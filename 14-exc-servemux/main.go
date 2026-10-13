package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("GET /users/{id}", usersHandler)
	mux.HandleFunc("GET /files/{path...}", filesHandler)
	mux.HandleFunc("GET /files/{$}", filesIndexHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func filesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("file index"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "user id=%d", id)
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	pathFile := r.PathValue("path")
	w.Write([]byte(pathFile))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/about" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Write([]byte("About page"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/health" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Write([]byte("ok"))
}
