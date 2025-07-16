package main

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
}

type CreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUser(email string, firstName string, lastName string, password string) (*User, error) {
	id := uuid.New()

	pwHashed, err := HashPassword(password)

	return &User{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  pwHashed,
	}, err
}

func CreateUserFromJSON(body io.ReadCloser) (*User, error) {
	createRequest := new(CreateUserRequest)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(createRequest)
	if err != nil {
		return &User{}, nil
	}

	user, err := NewUser(createRequest.Email, createRequest.FirstName, createRequest.LastName, createRequest.Password)

	return user, err
}
