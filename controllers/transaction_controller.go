package controllers

import (
	"catering/config"
	"catering/models"
	"catering/payment"
	"net/http"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/v4"
)

func CreateTransaction(e echo.Context) error {

	paket := models.Paket{}
	user := models.User{}
	regent := models.Regency{}
	trans := models.Transaction{}
	// order.Status = "ordered"
	e.Bind(&trans)

	config.DB.Where("id=?", &trans.PaketID).Find(&paket)
	err := config.DB.Model(&paket).Association("transaction").Find(&paket.Transaction).Error

	config.DB.Where("id=?", &trans.UserID).Find(&user)
	getErrors := config.DB.Model(&user).Association("transaction").Find(&user.Transaction).Error

	config.DB.Where("id=?", &trans.RegencyID).Find(&regent)
	//

	qty := trans.Quantity
	price := paket.Price
	shipping_cost := regent.ShippingCost
	total := (qty * price) + shipping_cost
	trans.Total = total
	config.DB.Save(&trans)

	if getErrors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error,
		})
	}

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error tahepa",
		})
	}

	paymentURL, err := payment.GetPaymentUrl(trans, user)
	trans.PaymentUrl = paymentURL
	config.DB.Save(&trans)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"order_id":           trans.ID,
		"Paket":              paket.Name,
		"quantity":           trans.Quantity,
		"total":              total,
		"lokasi pengantaran": regent.Name,
		"Nama Pemesa":        user.FullName,
		"paymnet url":        paymentURL,
	})
}
