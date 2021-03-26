package xchange_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	xchange "github.com/xchange-ee/wallet-sdk-go"
)

func TestCreateWallet(t *testing.T) {
	godotenv.Load(".env.test")
	client := xchange.NewClient(os.Getenv("XCHANGE_TOKEN"), os.Getenv("ENV"))
	walletRequest := xchange.WalletRequest{
		Currency: "KSOC",
	}
	walletResponse, errAPI, err := client.Wallet().Create(walletRequest)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if walletResponse == nil {
		t.Error("payResponse is null")
		return
	}

}

func TestCreatAddress(t *testing.T) {
	godotenv.Load(".env.test")
	client := xchange.NewClient(os.Getenv("XCHANGE_TOKEN"), os.Getenv("ENV"))
	walletRequest := xchange.WalletRequest{
		Currency: "KSOC",
	}
	walletResponse, errAPI, err := client.Wallet().Create(walletRequest)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if walletResponse == nil {
		t.Error("walletResponse is null")
		return
	}
	addressRequest := xchange.AddressRequest{
		WalletID: walletResponse.ID,
	}
	addressResponse, errAPI, err := client.Wallet().GetNewAddress(addressRequest)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if addressResponse == nil {
		t.Error("addressResponse is null")
		return
	}

}
