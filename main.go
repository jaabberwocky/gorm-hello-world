package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm-hello-world/models"
	"io"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB
var err error

func init() {
	// createschema is idempotent
	models.CreateSchema()

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database!")
	}
	fmt.Println("db connection established...")
}

func main() {

	// enable logging
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	// middleware for custom log format
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.GET("/products", getAll)
	router.POST("/products", postOne)
	router.PUT("products", putOne)
	router.GET("/products/:code", getOne)

	router.Run(":4531")

}

func getAll(c *gin.Context) {
	var allProducts []models.Product

	if err = db.Find(&allProducts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, allProducts)
	}
}

func getOne(c *gin.Context) {
	var product models.Product
	// obtained from router parameter
	code := c.Param("code")

	if err = db.Where("code = ?", code).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func postOne(c *gin.Context) {
	var product models.Product

	// check if valid JSON
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	// check if db insert has any error
	if err := db.Create(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": "code already exists in DB"})
		return
	}
	c.JSON(200, gin.H{
		"code":  product.Code,
		"price": product.Price})

}

func putOne(c *gin.Context) {
	var product models.Product
	var recordFound bool

	// check if valid JSON
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "invalid input; not JSON"})
		return
	}

	if product.Code == "" {
		c.JSON(400, gin.H{"error": "invalid input; code is empty"})
		return
	}

	u := models.Product{}
	db.Where("code = ?", product.Code).First(&u)
	if u.Code == "" {
		fmt.Println("record not found, will insert...")
	} else {
		recordFound = true
	}

	if err := db.Save(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": "db error"})
		return
	}
	c.JSON(200, gin.H{
		"updated": strconv.FormatBool(recordFound),
		"code":    product.Code,
		"price":   product.Price})

}
