package main

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/payment-management-system/db"
	"github.com/lucass-segura/payment-management-system/routes"
)

func main() {
	loadEnv()
	db.DBconnection()

	e := echo.New()
	e.GET("/", routes.IndexHandle)

	//Users

	e.Logger.Fatal(e.Start(":8001"))

}

func loadEnv() {
	envFile := filepath.Join("db", ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Error loading .env file")
	}
}
