# Go CRUD API

## Description
Simple Go example for me to get better acquainted with Go, particularly the Gin framework and the Gorm object relational-mapper.

The `init()` function creates a `sqlite3` database with randomly generated prices for 1000 products. Each product has a code `L+<number>` (e.g. `L145`), and is unique.

## How to install
`go run main.go` to run it and navigate to `localhost:4531` to see the API. Logging is enabled also to `gin.log` in the project directory.

*Optional*:
`go get github.com/codegangsta/gin` to get live gin-reloading. Run it using `gin -a 4531 main.go`. Any changes will be automatically detected for much faster development.

## Routes
*`/products/` : shows all products and prices
*`/products/<code>`: shows price for given product code
