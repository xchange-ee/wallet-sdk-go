package xchange

import (
	"encoding/json"
	"fmt"
	"time"
)

type WithdrawItem struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Address       string    `json:"address"`
	Currency      string    `json:"currency"`
	CompanyID     int64     `gorm:"column:company_id" json:"-"`
	WalletID      int64     `json:"wallet_id"`
	Amount        string    `gorm:"column:amount"`
	Fee           string    `gorm:"column:fee"`
	Total         string    `gorm:"column:total"`
	Txid          string    `json:"txid"`
	Status        string    `json:"status"`
	Blockhash     string    `json:"blockhash"`
	ErrorMessage  string    `json:"error_message"`
	Confirmations int32     `json:"confirmations"`
	Blockchain    string    `json:"-"`
	IsToken       bool
}

type WithdrawPagesQuery struct {
	PerPage int              `json:"per_page"`
	Page    int              `json:"page"`
	Query   *WithdrawQuery   `json:"query"`
	OrderBy *WithdrawOrderBy `json:"order_by"`
}

type WithdrawQuery struct {
	To         int64  `json:"to"`
	From       int64  `json:"from"`
	Currency   string `json:"currency"`
	WalletID   int64  `json:"wallet_id"`
	Address    string `json:"address"`
	Status     string `json:"status"`
	Txid       string `json:"txid"`
	Blockchain string `json:"-"`
	IsToken    bool
}
type WithdrawOrderBy struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type WithdrawPagesResponse struct {
	Data  []*WithdrawItem `json:"data"`
	Page  int             `json:"page"`
	Pages int             `json:"pages"`
	Total int             `json:"total"`
}

type CreateWithdraw struct {
	WalletID int64   `json:"wallet_id" validate:"required"`
	Address  string  `json:"address" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
}

// Withdraw is a structure manager all about Withdraw
type Withdraw struct {
	client *Xchange
}

//Withdraw - Instance de Withdraw
func (c *Xchange) Withdraw() *Withdraw {
	return &Withdraw{client: c}
}

func (x *Withdraw) List(req WithdrawPagesQuery) (*WithdrawPagesResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *WithdrawPagesResponse
	err, errAPI := x.client.Request("POST", "/api/withdraws", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (x *Withdraw) Create(req CreateWithdraw) (*WithdrawItem, *Error, error) {
	data, _ := json.Marshal(req)
	var response *WithdrawItem
	err, errAPI := x.client.Request("POST", "/api/withdraw", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (x *Withdraw) Get(id int32) (*WithdrawItem, *Error, error) {
	var response *WithdrawItem
	err, errAPI := x.client.Request("POST", fmt.Sprintf("/api/withdraw/%d", id), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
