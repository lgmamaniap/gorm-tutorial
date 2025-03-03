package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("./01-quick-start/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // Find product with integer primary key and save it to product
	fmt.Println("Find by id: ", product)
	db.First(&product, "code  = ?", "D42") // find product with code D42
	fmt.Println("Find by code: ", product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Println("Update product's price: ", product)

	// Update - update multiple fields
	// first way
	db.Model(&product).Updates(Product{Price: 300, Code: "F42"})
	fmt.Println("Update using struct: ", product)

	// second way
	db.Model(&product).Updates(map[string]any{"Price": 400, "Code": "G42"})
	fmt.Println("Update using map: ", product)

	// Delete - delete product
	db.Delete(&product, 1)

	// db.Unscoped().Delete(&product, 1) // delete product permanently
}
