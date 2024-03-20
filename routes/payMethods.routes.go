package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/payment-management-system/db"
	"github.com/lucass-segura/payment-management-system/models"
)

func PostPayMethods(c echo.Context) error {
	var payMethod models.PayMethod

	json.NewDecoder(c.Request().Body).Decode(&payMethod)
	createPayMethod := db.DB.Create(&payMethod)
	err := createPayMethod.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error:": "Error creating rol in database"})
	}

	fmt.Println(payMethod.PayMethod, payMethod.ID)
	return c.JSON(http.StatusCreated, payMethod)
}

func GetPayMethods(c echo.Context) error {
	var payMethods []models.PayMethod
	db.DB.Find(&payMethods)

	return json.NewEncoder(c.Response()).Encode(&payMethods)
}
