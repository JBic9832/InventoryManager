package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	ListenAddr string
	Storage    *sql.DB
}

type ServerError struct {
	Error string
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func EncodeJSONIntoResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func makeHandlerFunc(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			errorMessage := ServerError{
				Error: err.Error(),
			}
			EncodeJSONIntoResponse(w, http.StatusBadRequest, errorMessage)
		}
	}
}

func NewServer(listenAddr string) *Server {
	db, err := sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		ListenAddr: listenAddr,
		Storage:    db,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/", s.Hello)

	log.Printf("Server is live at http://localhost%s", s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, r)
}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
