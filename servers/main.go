package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

// Структура сервера
type Server struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	CPU    int    `json:"cpu"`
	RAM    int    `json:"ram"`
	Region string `json:"region"`
}

var (
	servers  []Server
	idSeq    int
	mutex    sync.Mutex
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/servers", listServers).Methods("GET")
	r.HandleFunc("/servers", addServer).Methods("POST")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func listServers(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(servers)
}

func addServer(w http.ResponseWriter, r *http.Request) {
	var s Server
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	idSeq++
	s.ID = idSeq
	servers = append(servers, s)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
