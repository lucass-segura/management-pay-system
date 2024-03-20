package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/payment-management-system/db"
	"github.com/lucass-segura/payment-management-system/models"
)

func CreateRolHandle(c echo.Context) error {
	var rol models.Role
	fmt.Println(c.Request().Body)

	json.NewDecoder(c.Request().Body).Decode(&rol)
	createdRol := db.DB.Create(&rol)
	err := createdRol.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error:": "Error creating rol in database"})
	}
	return c.JSON(http.StatusCreated, rol)
}
