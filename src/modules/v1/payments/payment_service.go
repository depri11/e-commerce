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

// func GenerateSnapReq(transaction *models.Transaction, user *models.User) *snap.Request {
// 	var s = snap.Client
// 	s.New("YOUR-SERVER-KEY", midtrans.Sandbox)

// 	snapReq := &snap.Request{
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  transaction.Code,
// 			GrossAmt: int64(transaction.Amount),
// 		},
// 		CustomerDetail: &midtrans.CustomerDetails{
// 			FName: user.Name,
// 			Email: user.Email,
// 		},
// 	}

// 	return snapReq
// }

func (s *service) GetPaymentURL(transaction *models.Transaction, user *models.User) (string, error) {
	id := transaction.ID.Hex()

	midclient := midtrans.NewClient()
	midclient.ClientKey = "SB-Mid-client-Fg55R6OSZynaFTNA"
	midclient.ServerKey = "SB-Mid-server-Dc-ShUJ8AJYb9EvWzoVZKCq0"
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
