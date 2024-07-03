package config

import (
	"github.com/veritrans/go-midtrans"
)

var MidtransClient midtrans.Client

func InitMidtrans() {
	MidtransClient = midtrans.NewClient()
	MidtransClient.ServerKey = "YOUR_SERVER_KEY"
	MidtransClient.ClientKey = "YOUR_CLIENT_KEY"
	MidtransClient.APIEnvType = midtrans.Sandbox // Use midtrans.Production for production environment
}
