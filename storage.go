package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Database struct {
	Store *sql.DB
}

func NewDatabase() *Database {
	db, err := sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Store: db,
	}
}

func (db *Database) AddUserToDatabase(user *User) error {
	query := `
	INSERT INTO Users(UserID, Email, FirstName, LastName, Password)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err := db.Store.Exec(query, user.ID.String(), user.Email, user.FirstName, user.LastName, user.Password)
	return err
}
