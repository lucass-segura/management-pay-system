package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/payment-management-system/db"
	"github.com/lucass-segura/payment-management-system/models"
)

func PostUserHandler(c echo.Context) error {
	var user models.User
	fmt.Println(c.Request().Body)

	json.NewDecoder(c.Request().Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating user in database."})
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUsersHandler(c echo.Context) error {
	var users []models.User
	db.DB.Find(&users)

	return json.NewEncoder(c.Response()).Encode(&users)
}

func GetUserHandler(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	userID, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid user id.")
	}

	db.DB.First(&user, userID)

	if user.ID == 0 {
		c.String(http.StatusNotFound, "Not found user id.")
	}

	return json.NewEncoder(c.Response()).Encode(&user)
}

func DeleteUserHandler(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	UserID, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id.")
	}

	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found user id.")
	}

	db.DB.Unscoped().Delete(&user, UserID)

	return c.String(http.StatusAccepted, "User deteled.")
}
