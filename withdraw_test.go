package xchange_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	xchange "github.com/xchange-ee/wallet-sdk-go"
)

func TestListWithdraw(t *testing.T) {
	godotenv.Load(".env.test")
	client := xchange.NewClient(os.Getenv("XCHANGE_TOKEN"), os.Getenv("ENV"))
	WithdrawPagesQuery := xchange.WithdrawPagesQuery{
		PerPage: 10,
		Page:    1,
		Query: &xchange.WithdrawQuery{
			Blockchain: "BSC",
		},
		OrderBy: &xchange.WithdrawOrderBy{},
	}
	WithdrawPagesResponse, errAPI, err := client.Withdraw().List(WithdrawPagesQuery)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if WithdrawPagesResponse == nil {
		t.Error("payResponse is null")
		return
	}

}

func TestCreate(t *testing.T) {
	godotenv.Load(".env.test")
	client := xchange.NewClient(os.Getenv("XCHANGE_TOKEN"), os.Getenv("ENV"))
	WithdrawPagesQuery := xchange.CreateWithdraw{
		WalletID: 6,
		Address:  "0x6732D4B302795dae4dB236C7dB29592fc4268AB9",
		Amount:   1,
	}
	WithdrawPagesResponse, errAPI, err := client.Withdraw().Create(WithdrawPagesQuery)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if WithdrawPagesResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
