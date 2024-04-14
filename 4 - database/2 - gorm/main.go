package main

import (
	"fmt"

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
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.13,
	// })

	//create batch
	// products := []Product{
	// 	{Name: "Desktop", Price: 6000.00},
	// 	{Name: "Mouse", Price: 250.00},
	// 	{Name: "Keyboard", Price: 420.10},
	// }
	// db.Create(&products)

	//select one
	// var product Product
	// db.First(&product, 3)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Keyboard")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//select with limit
	// var products []Product
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//select with offset
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//where
	// var products []Product
	// db.Where("price > ?", 1000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//like
	var products []Product
	db.Where("name like ?", "%k%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
