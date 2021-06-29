package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetRegency(e echo.Context) error {
	regencys := []models.Regency{}
	err := config.DB.Find(&regencys).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	response := helpers.Apiresponse("list of regency", http.StatusOK, "success", regencys)
	return e.JSON(http.StatusOK, response)

}

func CreateRegency(e echo.Context) error {

	regency := models.Regency{}
	e.Bind(&regency)
	err := config.DB.Save(&regency).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("Regency has been created", http.StatusOK, "success", regency)
	return e.JSON(http.StatusOK, response)
}

func ShowRegency(e echo.Context) error {
	regency := models.Regency{}
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id=?", id).Find(&regency)
	err := config.DB.Model(&regency).Association("transaction").Find(&regency.Transaction).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	response := helpers.Apiresponse("regency found", http.StatusOK, "success", regency)
	return e.JSON(http.StatusOK, response)

}

func UpdateRegency(e echo.Context) error {
	regency := models.Regency{}
	id, _ := strconv.Atoi(e.Param("id"))
	// name := e.FormValue("name")
	// cost, _ := strconv.Atoi(e.FormValue("shipping_cost"))
	err := config.DB.Find(&regency, id).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	e.Bind(&regency)
	errors := config.DB.Save(&regency).Error
	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("Regency has been updated", http.StatusOK, "success", regency)
	return e.JSON(http.StatusOK, response)

}

//can't delete regency cause it permanent
func DeleteRegency(e echo.Context) error {

	regency := models.Regency{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&regency).Error
	config.DB.Delete(&regency)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("Regency has been deleted", http.StatusOK, "success", regency)
	return e.JSON(http.StatusOK, response)
}
