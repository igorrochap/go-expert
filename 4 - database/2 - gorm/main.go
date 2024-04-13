package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.13,
	})

	//create batch
	products := []Product{
		{Name: "Desktop", Price: 6000.00},
		{Name: "Mouse", Price: 250.00},
		{Name: "Keyboard", Price: 420.10},
	}
	db.Create(&products)
}
