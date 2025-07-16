package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	ListenAddr string
	Storage    *Database
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

	db := NewDatabase()

	return &Server{
		ListenAddr: listenAddr,
		Storage:    db,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/user", makeHandlerFunc(s.AddUserToDatabase))

	log.Printf("Server is live at http://localhost%s", s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, r)
}
