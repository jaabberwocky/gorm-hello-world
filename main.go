package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm-hello-world/models"
	"io"
	"os"
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

		// your custom format
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
	router.GET("/products/:code", getOne)
	router.POST("/products/add", postOne)

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
		c.String(404, "Not found!")
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
		c.JSON(400, gin.H{"error": "error occured with db creation"})
		return
	}
	c.JSON(200, gin.H{
		"code":  product.Code,
		"price": product.Price})
		}

}
