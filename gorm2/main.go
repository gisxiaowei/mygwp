package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var Db *gorm.DB

// connect to the Db
func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:chinaren@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	Db.AutoMigrate(&Product{})
}

func main() {
	// Create
	Db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	Db.First(&product, 1) // find product with id 1
	fmt.Println(product)
	Db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println(product)

	// Update - update product's price to 2000
	Db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	Db.Delete(&product)
}
