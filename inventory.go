package main

type Product struct {
	UPC                int64   `json:"upc"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	DepartmentID       int     `json:"department_id"`
	Price              float32 `json:"price"`
}

type Department struct {
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
}

type ProductEntry struct {
	UPC      int64 `json:"upc"`
	Quantity int   `json:"quantity"`
}

func NewDepartment(departmentId int, departmentName string) *Department {
	return &Department{
		DepartmentID:   departmentId,
		DepartmentName: departmentName,
	}
}

func NewProduct(upc int64, name string, description string,
	departmentId int, price float32) *Product {
	return &Product{
		UPC:                upc,
		ProductName:        name,
		ProductDescription: description,
		DepartmentID:       departmentId,
		Price:              price,
	}
}
