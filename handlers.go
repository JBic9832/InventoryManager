package main

import (
	"net/http"
)

func (s *Server) AddUserToDatabase(w http.ResponseWriter, r *http.Request) error {
	user, err := CreateUserFromJSON(r.Body)
	if err != nil {
		return err
	}

	return s.Storage.AddUserToDatabase(user)
}
