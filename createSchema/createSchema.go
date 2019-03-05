package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"unique_index; primary_key"`
	Price uint
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("uh-oh")
	}
	fmt.Printf("Current working directory %s", cwd)
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// REMOVE IF NECC: wipe DB
	db.Exec("DROP TABLE IF EXISTS products;")

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	p := Product{Code: "L44499", Price: 2500}
	db.Create(&p)

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// sleep then update
	time.Sleep(5 * time.Second)
	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

}
