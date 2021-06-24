package controllers

import (
	"catering/config"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetPakets(e echo.Context) error {
	var Pakets []models.Paket
	err := config.DB.Find(&Pakets).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success",
		"data":    Pakets,
	})

}

func CreatePaket(e echo.Context) error {
	Paket := models.Paket{}
	e.Bind(&Paket)
	err := config.DB.Save(&Paket).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Paket created",
		"data":    Paket,
	})
}

func ShowPaket(e echo.Context) error {

	Paket := models.Paket{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&Paket, id).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Paket found",
		"data":    Paket,
	})
}

func UpdatePaket(e echo.Context) error {

	Paket := models.Paket{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&Paket, id).Error
	e.Bind(&Paket)
	config.DB.Save(&Paket)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Paket updated",
		"data":    Paket,
	})
}

func DeletePaket(e echo.Context) error {

	Paket := models.Paket{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Table("Pakets").Where("id=?", id).Delete(&Paket).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Paket Deleted",
		"data":    Paket,
	})
}
