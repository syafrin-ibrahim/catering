package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/models"
	_ "catering/payment"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/v4"
)

func CreateTransaction(e echo.Context) error {

	paket := models.Paket{}
	user := models.User{}
	regent := models.Regency{}
	trans := models.Transaction{}

	// regencyId, _ := strconv.Atoi(e.FormValue("regency_id"))
	// userId, _ := strconv.Atoi(e.FormValue("user_id"))
	// paketId, _ := strconv.Atoi(e.FormValue("paket_id"))
	// qty, _ := strconv.Atoi(e.FormValue("quantity"))
	// location := e.FormValue("location")
	// deliverAt := e.FormValue("deliver_time")

	e.Bind(&trans)

	err := config.DB.Where("id=?", trans.PaketID).Find(&paket).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "Error query paket ",
			"state":  err.Error(),
		})
	}

	config.DB.Where("id=?", trans.UserID).Find(&user)

	config.DB.Where("id=?", trans.RegencyID).Find(&regent)

	qty := trans.Quantity
	price := paket.Price
	shipping_cost := regent.ShippingCost
	total := (qty * price) + shipping_cost

	newTrans := models.Transaction{
		UserID:      trans.UserID,
		PaketID:     trans.PaketID,
		Quantity:    qty,
		Total:       total,
		Location:    trans.Location,
		RegencyID:   trans.RegencyID,
		Status:      "pending",
		DeliverTime: trans.DeliverTime,
		Note:        trans.Note,
	}
	//simpan transaksi
	nextError := config.DB.Save(&newTrans).Error
	if nextError != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error during create transaction",
			"message": nextError.Error,
		})
	}

	//set payment url
	// paymentURL, getErr := payment.GetPaymentUrl(newTrans, user)

	// if getErr != nil {
	// 	return e.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  "error during get payment url",
	// 		"message": getErr.Error,
	// 	})
	// }

	// // update transaski set payment url
	// newTrans.PaymentUrl = paymentURL

	// getErrors := config.DB.Save(&newTrans).Error
	// if getErrors != nil {
	// 	return e.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  "error during create transaction",
	// 		"message": getErrors.Error,
	// 	})
	// }

	transactionResponse := models.TransactionResponse{
		TransID:       newTrans.ID,
		PaketName:     paket.Name,
		Quantity:      newTrans.Quantity,
		Total:         newTrans.Total,
		Location:      newTrans.Location,
		RegentName:    regent.Name,
		CustomerName:  user.FullName,
		DeliveredTime: newTrans.DeliverTime,
		PaymentURL:    "",
		Note:          newTrans.Note,
	}

	response := helpers.Apiresponse("Transaction Created", http.StatusOK, "success", transactionResponse)
	return e.JSON(http.StatusOK, response)
	// return e.JSON(http.StatusOK, map[string]interface{}{
	// 	"status":  http.StatusOK,
	// 	"message": "Transaction Created",
	// 	"data":    transactionResponse,
	// })
}

func GetTransaction(e echo.Context) error {
	transactions := []models.Transaction{}
	// err := config.DB.Find(&transactions).Error
	err := config.DB.Preload("Paket").Preload("User").Preload("Regency").Find(&transactions).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var formatTransResponse []models.TransactionResponse
	for _, transaction := range transactions {
		transactionResponse := models.TransactionResponse{
			TransID:       transaction.ID,
			PaketName:     transaction.Paket.Name,
			Quantity:      transaction.Quantity,
			Total:         transaction.Total,
			Location:      transaction.Location,
			RegentName:    transaction.Regency.Name,
			CustomerName:  transaction.User.FullName,
			DeliveredTime: transaction.DeliverTime,
			PaymentURL:    "",
			Note:          transaction.Note,
		}

		formatTransResponse = append(formatTransResponse, transactionResponse)
	}

	response := helpers.Apiresponse("list of transaction", http.StatusOK, "success", formatTransResponse)
	return e.JSON(http.StatusOK, response)

}

func GetTransactionByUserId(e echo.Context) error {
	var transactions []models.Transaction
	// currentUser := e.Get("currentUser").(models.User)
	// userId := currentUser.ID
	userId := 1
	err := config.DB.Preload("Paket").Preload("User").Preload("Regency").Where("user_id=?", userId).Find(&transactions).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var formatTransResponse []models.TransactionResponse
	for _, transaction := range transactions {
		transactionResponse := models.TransactionResponse{
			TransID:       transaction.ID,
			PaketName:     transaction.Paket.Name,
			Quantity:      transaction.Quantity,
			Total:         transaction.Total,
			Location:      transaction.Location,
			RegentName:    transaction.Regency.Name,
			CustomerName:  transaction.User.FullName,
			DeliveredTime: transaction.DeliverTime,
			PaymentURL:    "",
			Note:          transaction.Note,
		}

		formatTransResponse = append(formatTransResponse, transactionResponse)
	}

	response := helpers.Apiresponse("list of transaction", http.StatusOK, "success", formatTransResponse)
	return e.JSON(http.StatusOK, response)
}

func GetNotification(e echo.Context) error {

	notif := models.MidtransNotification{}
	e.Bind(&notif)

	transaction := models.Transaction{}
	transaction_id, _ := strconv.Atoi(notif.OrderID)

	config.DB.Where("id=?", transaction_id).First(&transaction)

	if notif.PaymentType == "credit_card" && notif.TransactionStatus == "capture" && notif.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if notif.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if notif.TransactionStatus == "deny" || notif.TransactionStatus == "expire" || notif.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	err := config.DB.Save(&transaction).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Oppss.... eror midtrans",
		})
	}
	return e.JSON(http.StatusOK, notif)
}
