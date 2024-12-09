package db

import (
	"fmt"
	"prd_mngt/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDatabase() {
	// Load environment variables
	utils.LoadEnv()

	// Database URL without SSL
	dbURL := "postgresql://postgres:5432@localhost:5432/prd_mngt?sslmode=require"

	// Alternatively, if you want to use an env variable, uncomment the line below
	//dbURL := utils.GetEnv("DATABASE_URL")

	// Open the database connection
	var err error
	DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Connected to Postgres")
	}

	// Auto-migrate the schema for User and Product models
	DB.AutoMigrate(&User{}, &Product{})
}

// User model
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product model
type Product struct {
	ID                      int      `json:"id"`
	UserID                  int      `json:"user_id"`
	ProductName             string   `json:"product_name"`
	ProductDescription      string   `json:"product_description"`
	ProductImages           []string `json:"product_images" gorm:"type:json"`
	ProductPrice            float64  `json:"product_price"`
	CompressedProductImages []byte   `json:"compressed_product_images"`
}
