package config

import (
	"log"
	"management-warehouse/category"
	"management-warehouse/mutation"
	"management-warehouse/product"
	"management-warehouse/stockadjustment"
	"management-warehouse/supplier"
	"management-warehouse/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres dbname=warehouse port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// dsn := os.Getenv("DATABASE_URI")
	if dsn == "" {
		log.Fatal("DATABASE_URI environment variable is not set")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = database.AutoMigrate(&user.User{}, &supplier.Supplier{}, &product.Product{}, &category.Category{}, &mutation.Mutation{}, &stockadjustment.StockAdjustment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	DB = database
	log.Println("Database connection established")
}
