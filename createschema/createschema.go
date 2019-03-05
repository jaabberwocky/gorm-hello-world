package createschema

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

type Product struct {
	Code  string `gorm:"unique_index; primary_key" json:"code"`
	Price uint   `json:"price"`
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

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	p := Product{Code: "L44499", Price: 2500}
	db.Create(&p)

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println("Created test.db...")
}
