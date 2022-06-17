package payment

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/veritrans/go-midtrans"
)

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction models.Payment, user models.User) (string, error) {
	id := transaction.ID

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-client-Fg55R6OSZynaFTNA"
	midclient.ClientKey = "SB-Mid-server-Dc-ShUJ8AJYb9EvWzoVZKCq0"
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	chargeReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  id,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := coreGateway.GetToken(chargeReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
