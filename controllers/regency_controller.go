package controllers

import (
	"catering/config"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetRegency(e echo.Context) error {
	Regencys := []models.Regency{}
	err := config.DB.Find(&Regencys).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success",
		"data":    Regencys,
	})

}

func CreateRegency(e echo.Context) error {
	Regency := models.Regency{}
	e.Bind(&Regency)
	err := config.DB.Save(&Regency).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Regency created",
		"data":    Regency,
	})
}

func ShowRegency(e echo.Context) error {
	Regency := models.Regency{}
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id=?", id).Find(&Regency)
	err := config.DB.Model(&Regency).Association("transaction").Find(&Regency.Transaction).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Regency found",
		"data":    Regency,
	})

}

func UpdateRegency(e echo.Context) error {
	Regency := models.Regency{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&Regency, id).Error
	e.Bind(&Regency)
	config.DB.Save(&Regency)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Regency updates",
		"data":    Regency,
	})

}

//can't delete regency cause it permanent
// func DeleteRegency(e echo.Context) error {

// 	Regency := models.Regency{}
// 	id, _ := strconv.Atoi(e.Param("id"))
// 	err := config.DB.Where("id=?", id).Find(&Regency).Error
// 	config.DB.Delete(&Regency)

// 	if err != nil {
// 		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}

// 	return e.JSON(http.StatusOK, map[string]interface{}{
// 		"status":  http.StatusOK,
// 		"message": "Regency Deleted",
// 		"data":    Regency,
// 	})
// }
