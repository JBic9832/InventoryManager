package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	db, err := sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		DB: db,
	}

}

func AddProductToDB(s *sql.DB, p Product) error {
	_, err := s.Exec(`
		INSERT INTO Products(UPC, ProductName, ProductDescription, DepartmentID, Price)
		VALUES (?, ?, ?, ?, ?)`, p.UPC, p.ProductName, p.ProductDescription, p.DepartmentID, p.Price)
	return err
}
