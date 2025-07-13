package main

import "database/sql"

func AddProductToDB(s *sql.DB, p Product) error {
	_, err := s.Exec(`
		INSERT INTO Products(UPC, ProductName, ProductDescription, DepartmentID, Price)
		VALUES (?, ?, ?, ?, ?)`, p.UPC, p.ProductName, p.ProductDescription, p.DepartmentID, p.Price)
	return err
}
