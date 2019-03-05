package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User test if we can read users table
type User struct {
	ID    int
	fname string
	lname string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	u := User{}
	if err := db.Where("fname = ?", "james").First(&u); err != nil {
		panic("could not find query")
	}
	fmt.Printf("%+v", u)

}
