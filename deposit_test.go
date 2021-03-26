package xchange_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	xchange "github.com/xchange-ee/wallet-sdk-go"
)

func TestListDeposits(t *testing.T) {
	godotenv.Load(".env.test")
	client := xchange.NewClient(os.Getenv("XCHANGE_TOKEN"), os.Getenv("ENV"))
	depositPagesQuery := xchange.DepositPagesQuery{
		PerPage: 10,
		Page:    1,
		Query: &xchange.DepositQuery{
			Currency: "KSOC",
		},
		OrderBy: &xchange.DepositOrderBy{},
	}
	depositPagesResponse, errAPI, err := client.Deposit().List(depositPagesQuery)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if depositPagesResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
