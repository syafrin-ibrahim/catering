package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetCategories(e echo.Context) error {
	var categories []models.Category
	err := config.DB.Find(&categories).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("list of category", http.StatusOK, "success", categories)
	return e.JSON(http.StatusOK, response)
}

func CreateCategory(e echo.Context) error {
	category := models.Category{}
	e.Bind(&category)
	err := config.DB.Save(&category).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("category has been created", http.StatusOK, "success", category)
	return e.JSON(http.StatusOK, response)
}

func ShowCategory(e echo.Context) error {

	category := models.Category{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&category, id).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("category found", http.StatusOK, "success", category)
	return e.JSON(http.StatusOK, response)
}

func UpdateCategory(e echo.Context) error {

	category := models.Category{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&category, id).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	e.Bind(&category)
	getErrors := config.DB.Save(&category).Error
	if getErrors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("category has been updated", http.StatusOK, "success", category)
	return e.JSON(http.StatusOK, response)
}

func DeleteCategory(e echo.Context) error {

	category := models.Category{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&category).Error
	config.DB.Delete(&category)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("category has been deleted", http.StatusOK, "success", category)
	return e.JSON(http.StatusOK, response)
}
