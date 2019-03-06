package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // used to create schema for sqlite3
	"log"
	"math/rand"
	"os"
	"strconv"
)

// Product contains Code and Price
type Product struct {
	Code  string `gorm:"unique_index; primary_key" json:"code"`
	Price int    `json:"price,string"`
}

// CreateSchema creates schema
func CreateSchema() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("uh-oh")
	}
	fmt.Printf("Current working directory %s\n", cwd)
	fmt.Println("Connecting to db...")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// REMOVE IF NECC: wipes DB
	fmt.Println("Dropping existing table...")
	db.Exec("DROP TABLE IF EXISTS products;")

	// Migrate the schema
	db.AutoMigrate(&Product{})

	fmt.Println("Creating database...")
	for i := 0; i < 1000; i++ {
		var p Product
		p.Code = "L" + strconv.Itoa(i)
		p.Price = rand.Intn(1000) + 250
		db.Create(&p)
	}

	fmt.Println("Created test.db...")
}
