package payment

import (
	"strconv"

	"fiber-e-commerce-system-API/domain/models"
	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	client midtrans.Client
}

type Service interface {
	GetPaymentURL(transaction Transaction, user models.User) (string, error)
}

func NewService(client midtrans.Client) *service {
	return &service{client: client}
}

func (s *service) GetPaymentURL(transaction Transaction, user models.User) (string, error) {
	snapGateway := midtrans.SnapGateway{
		Client: s.client,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Total),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
