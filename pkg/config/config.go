package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// load the env
func LoadEnv() {
	err := godotenv.Load("C:/Users/USER/Desktop/go-backend-project/draw-app-js-go/.env")
	if err != nil {
		log.Fatal("Error while loading env file")
	}
}

// esess the thhis env every ware
func GetEnv(key string) string {
	return os.Getenv(key)
}

// conect to db
func ConectDb() {
	LoadEnv()

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmodel=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to conect with database", err)
	}

	fmt.Println("database conected")
}

func GetDb() *gorm.DB {
	return DB
}
