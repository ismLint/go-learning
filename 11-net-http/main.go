package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tasks  []Task
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", statusHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskByIDHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет!"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("ok"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "method=%s path=%s", r.Method, r.URL.Path)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "гость"
	}

	fmt.Fprintf(w, "Привет, %s!", name)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Write(body)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	var task Task

	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&task)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		task.ID = nextID
		nextID++
		defer mu.Unlock()
		tasks = append(tasks, task)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func taskByIDHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet && r.Method != http.MethodPut &&
		r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	stringByID := strings.TrimPrefix(r.URL.Path, "/tasks/")
	taskID, err := strconv.Atoi(stringByID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var found bool
	var foundTask Task
	var updateTask Task

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()

		for index := range tasks {
			if tasks[index].ID == taskID {
				found = true
				foundTask = tasks[index]
				break
			}
		}

		if found {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(foundTask)
		} else {
			http.Error(w, "Task Not Found", http.StatusNotFound)
			return
		}
	case http.MethodPut:
		defer r.Body.Close()

		var foundIndex int

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateTask)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()

		for index := range tasks {
			if tasks[index].ID == taskID {
				found = true
				foundIndex = index
				break
			}
		}

		if !found {
			http.Error(w, "Task Not Found", http.StatusNotFound)
			return
		} else {
			updateTask.ID = taskID
			tasks[foundIndex] = updateTask

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateTask)
		}
	case http.MethodDelete:
		mu.Lock()
		var indexToRemove int

		for index := range tasks {
			if tasks[index].ID == taskID {
				found = true
				indexToRemove = index
				break
			}
		}

		if !found {
			mu.Unlock()
			http.Error(w, "Task Not Found", http.StatusNotFound)
			return
		}

		tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)
		mu.Unlock()

		w.WriteHeader(http.StatusNoContent)
	}
}
