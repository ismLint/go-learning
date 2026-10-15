package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status      int
	bytes       int
	wroteHeader bool
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /{$}", helloHandler)
	mux.HandleFunc("GET /files/{path...}", filesHandler)
	mux.HandleFunc("GET /files/{$}", filesIndexHandler)
	mux.HandleFunc("GET /users/{id}", usersHandler)
	mux.HandleFunc("GET /temp-status", tempStatusHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           statusLoggingMiddleware(loggingMiddleware(mux)),
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

func (sw *statusWriter) WriteHeader(status int) {
	if sw.wroteHeader {
		return
	}

	sw.wroteHeader = true
	sw.ResponseWriter.WriteHeader(status)
}

func (sw *statusWriter) Write(data []byte) (int, error) {
	if !sw.wroteHeader {
		sw.WriteHeader(http.StatusOK)
	}

	n, err := sw.ResponseWriter.Write(data)
	sw.bytes += n

	return n, err
}

func tempStatusHandler(w http.ResponseWriter, r *http.Request) {
	sw := statusWriter{ResponseWriter: w}

	sw.tempHandler()
}

func (sw *statusWriter) tempHandler() {
	sw.WriteHeader(201) //change statusWriter.status
	sw.WriteHeader(400)

	sw.Write([]byte("created"))
}

func statusLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sw := &statusWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(sw, r)

		log.Printf("%s %s status=%d bytes=%d", r.Method, r.URL.Path, sw.status, sw.bytes)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s in %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello page"))
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	pathFile := r.PathValue("path")
	w.Write([]byte(pathFile))
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
