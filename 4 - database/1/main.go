package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func newProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := newProduct("Notebook", 4699.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 4000.99
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	result, err := findByID(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product %v has the price of %.2f\n", result.Name, result.Price)
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func findByID(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
