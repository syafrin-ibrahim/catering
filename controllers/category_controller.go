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
	//err := config.DB.Find(&categories).Error
	err := config.DB.Preload("Paket").Find(&categories).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var formatCategories []models.CategoryResponse

	for _, category := range categories {
		categoriResponse := models.CategoryResponse{
			ID:    category.ID,
			Name:  category.Name,
			Paket: category.Paket,
		}
		formatCategories = append(formatCategories, categoriResponse)
	}

	response := helpers.Apiresponse("list of category", http.StatusOK, "success", formatCategories)
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
	err := config.DB.Preload("Paket").Find(&category, id).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var paketFormat []models.DetailPaketResponse
	for _, paket := range category.Paket {
		images, _ := FindImage(paket.ID)
		paketResponse := models.DetailPaketResponse{
			ID:          paket.ID,
			Name:        paket.Name,
			Price:       paket.Price,
			Description: paket.Description,
			Discount:    paket.Discount,
			Image:       images,
		}

		paketFormat = append(paketFormat, paketResponse)
	}

	response := helpers.Apiresponse("category "+category.Name, http.StatusOK, "success", paketFormat)
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

func FindImage(id int) ([]models.Image, error) {
	var image []models.Image
	err := config.DB.Where("id=?", id).Find(&image).Error
	if err != nil {
		return image, nil
	}

	return image, nil
}
