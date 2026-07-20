package config

import (
	"fmt"

	"github.com/veritrans/go-midtrans"
)

func NewMidtransClient(cfg Config) (midtrans.Client, error) {
	client := midtrans.NewClient()
	client.ServerKey = cfg.MidtransServerKey
	client.ClientKey = cfg.MidtransClientKey

	switch cfg.MidtransEnv {
	case "sandbox":
		client.APIEnvType = midtrans.Sandbox
	case "production":
		client.APIEnvType = midtrans.Production
	default:
		return midtrans.Client{}, fmt.Errorf("invalid environment variable MIDTRANS_ENV")
	}

	return client, nil
}
