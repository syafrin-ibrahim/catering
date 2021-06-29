package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

var hostUrl = "http://127.0.0.1:8000/"

func CreateImage(e echo.Context) error {
	image := models.Image{}

	e.Bind(&image)
	var mainImage bool
	if image.IsMain == true {
		// _, err := makeImageNonPrimary(paketId)
		// if err != nil {

		// 	return e.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 		"message": "ooopss.... sql error",
		// 		"status":  err.Error,
		// 	})
		// }

		mainImage = true
	} else {
		mainImage = false
	}

	//err := config.DB.Where("id=?", paketId).First(&paket).Error
	findPaket, err := FindPaketById(image.PaketID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error sql",
			"error":   err.Error(),
		})
	}
	//paketName := paket.Name
	paketName := strings.Replace(findPaket.Name, " ", "_", -1)
	file, err := e.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filePath := "./uploads/" + paketName + "_" + file.Filename
	fileSrc := "uploads/" + paketName + "_" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	newImage := models.Image{
		PaketID:  findPaket.ID,
		FileName: fileSrc,
		IsMain:   mainImage,
	}

	errors := config.DB.Create(&newImage).Error

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "ooopss.... there is something error",
			"status":  err.Error,
		})
	}

	imageResponse := models.ImageResponse{
		PaketID:  newImage.PaketID,
		FileName: hostUrl + fileSrc,
		IsMain:   mainImage,
	}

	response := helpers.Apiresponse("Image Has Been Uploaded", http.StatusOK, "success", imageResponse)
	return e.JSON(http.StatusOK, response)
}

func makeImageNonPrimary(paketId int) (bool, error) {

	err := config.DB.Model(&models.Image{}).Where("id=?", paketId).Update("is_primary", false).Error

	if err != nil {
		return false, nil
	}

	return true, nil

}

func DeleteImage(e echo.Context) error {
	image := models.Image{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).First(&image).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	imageResource := image.FileName

	if FileExists(imageResource) {
		e := os.Remove(imageResource)
		if e != nil {
			log.Fatal(e)
		}
	}

	errors := config.DB.Table("images").Where("id=?", id).Delete(&image).Error

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Image Deleted",
	})
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FindPaketById(id int) (models.Paket, error) {
	var paket models.Paket
	err := config.DB.Where("id=?", id).Find(&paket).Error

	if err != nil {
		return paket, nil
	}

	return paket, nil

}
