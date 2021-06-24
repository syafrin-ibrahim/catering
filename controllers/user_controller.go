package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsers(e echo.Context) error {
	var Users []models.User
	err := config.DB.Find(&Users).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messsage": err.Error(),
		})

	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success",
		"data":    Users,
	})

}

func CreateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)
	hashPass, errors := helpers.HashPassword(user.Password)
	user.Password = string(hashPass)
	err := config.DB.Save(&user).Error

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "hashing password",
		})
	}

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "while saving obect user",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User created",
		"data":    user,
	})
}

func ShowUser(e echo.Context) error {
	User := models.User{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&User).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User found",
		"data":    User,
	})

}

func UpdateUser(e echo.Context) error {
	User := models.User{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&User, id).Error
	e.Bind(&User)
	config.DB.Save(&User)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User updates",
		"data":    User,
	})

}

func DeleteUser(e echo.Context) error {

	User := models.User{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&User).Error
	config.DB.Delete(&User)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User Deleted",
		"data":    User,
	})
}
