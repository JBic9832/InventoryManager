package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	ListenAddr string
	Storage    *sql.DB
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
