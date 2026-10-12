package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloHandle)
	http.HandleFunc("/echo", echoHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadHeaderTimeout: 5 * time.Second,  // ограничивает время на чтение заголовков запроса
		IdleTimeout:       60 * time.Second, // ограничиваем чтение всего запроса, вклчая тело
		ReadTimeout:       15 * time.Second, // задает предел времени на отпавку ответа
		WriteTimeout:      10 * time.Second, // ограничивает ожидание следующего запроса на keep-alive соединении

	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("server is running"))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) //ограничивает размер тела запроса 1 МБ

	defer r.Body.Close()

	var maxByteErr *http.MaxBytesError

	body, err := io.ReadAll(r.Body)
	if err != nil {
		if errors.As(err, &maxByteErr) {
			http.Error(w, "Content Too Large", http.StatusRequestEntityTooLarge)
			return
		} else {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}

	w.Write(body)
}
