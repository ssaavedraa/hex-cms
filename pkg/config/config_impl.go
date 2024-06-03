package config

import (
	"fmt"
	"log"
	"os"

	"hex/cms/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigImpl struct {}

func NewConfig () Config {
	return &ConfigImpl{}
}

func (c *ConfigImpl) LoadConfig () {
	env := c.GetEnv("ENVIRONMENT")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.GetEnv("DB_HOST"),
		c.GetEnv("DB_USER"),
		c.GetEnv("DB_PASSWORD"),
		c.GetEnv("DB_NAME"),
		c.GetEnv("DB_PORT"),
	)

	DbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = DbInstance

	err = DB.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Invoice{},
		&models.Shift{},
		&models.InvoiceItem{},
		&models.Product{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func (c *ConfigImpl) GetEnv (key string) string {
	return os.Getenv(key)
}