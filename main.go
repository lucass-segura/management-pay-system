package main

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/payment-management-system/db"
	"github.com/lucass-segura/payment-management-system/models"
	"github.com/lucass-segura/payment-management-system/routes"
)

func main() {
	loadEnv()
	db.DBconnection()

	db.DB.AutoMigrate(models.Role{})
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.PayMethod{})
	db.DB.AutoMigrate(models.Payment{})

	e := echo.New()
	e.GET("/", routes.IndexHandle)

	//Roles
	e.POST("/rol", routes.CreateRolHandle)
	e.GET("/rol", routes.GetRolesHandle)

	//Users
	e.POST("/users", routes.PostUserHandler)
	e.GET("/users", routes.GetUsersHandler)
	e.GET("/users/:id", routes.GetUserHandler)
	e.DELETE("/users/:id", routes.DeleteUserHandler)

	//Pay Methods
	e.POST("/methods", routes.PostPayMethods)
	e.GET("/methods", routes.GetPayMethods)

	e.Logger.Fatal(e.Start(":8001"))
}

func loadEnv() {
	envFile := filepath.Join("db", ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Error loading .env file")
	}
}
