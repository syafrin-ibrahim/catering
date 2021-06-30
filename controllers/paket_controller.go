package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetPakets(e echo.Context) error {
	var pakets []models.Paket
	//err := config.DB.Find(&pakets).Error
	err := config.DB.Preload("Image", "images.is_main=1").Find(&pakets).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var formatPaketResponse []models.PaketResponse
	for _, paket := range pakets {
		paketResponse := models.PaketResponse{
			ID:          paket.ID,
			Name:        paket.Name,
			Description: paket.Description,
			Price:       paket.Price,
			Discount:    paket.Discount,
			Image:       paket.Image[0].FileName,
		}

		formatPaketResponse = append(formatPaketResponse, paketResponse)
	}

	response := helpers.Apiresponse("list of paket", http.StatusOK, "success", formatPaketResponse)
	return e.JSON(http.StatusOK, response)

}

func CreatePaket(e echo.Context) error {
	paket := models.Paket{}
	e.Bind(&paket)
	err := config.DB.Save(&paket).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("paket created", http.StatusOK, "success", paket)
	return e.JSON(http.StatusOK, response)
}

func ShowPaket(e echo.Context) error {

	paket := models.Paket{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Preload("Image").Find(&paket, id).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// var paketIma []models.PaketResponse

	// for _, image := range models.ImageResponse{

	// }
	paketResponse := models.DetailPaketResponse{
		ID:          paket.ID,
		Name:        paket.Name,
		Price:       paket.Price,
		Description: paket.Description,
		Discount:    paket.Discount,
		Image:       paket.Image,
	}
	response := helpers.Apiresponse("paket "+paket.Name, http.StatusOK, "success", paketResponse)
	return e.JSON(http.StatusOK, response)
}

func UpdatePaket(e echo.Context) error {

	paket := models.Paket{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&paket, id).Error
	e.Bind(&paket)
	config.DB.Save(&paket)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("paket updated", http.StatusOK, "success", paket)
	return e.JSON(http.StatusOK, response)
}

func DeletePaket(e echo.Context) error {

	paket := models.Paket{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&paket).Error
	config.DB.Delete(&paket)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response := helpers.Apiresponse("paket deleted", http.StatusOK, "success", paket)
	return e.JSON(http.StatusOK, response)
}

func FindCategory(id int) (models.Category, error) {
	var category models.Category
	err := config.DB.Where("id=?", id).Find(&category).Error
	if err != nil {
		return category, nil
	}

	return category, nil
}
