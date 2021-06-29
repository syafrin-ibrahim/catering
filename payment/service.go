package payment

import (
	"catering/models"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

// type service struct {
// }

// type Service interface {
// 	GetPaymentUrl(trans models.Transaction, user models.User) (string, error)
// }

// func NewService() *service {
// 	return &service{}
// }

func GetPaymentUrl(trans models.Transaction, user models.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-217kGZdahU203WWQgde1jBxB"
	midclient.ClientKey = "SB-Mid-client-bQIgi-NpisnRZctv"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.FullName,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(trans.ID),
			GrossAmt: int64(trans.Total),
		},
	}

	respon, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "something wrong guysss.... midtrans, error get token", err
	}

	return respon.RedirectURL, nil

}
