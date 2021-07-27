package xchange

import (
	"encoding/json"
	"time"
)

type DepositItem struct {
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Address       string    `json:"address"`
	Amount        string    `gorm:"column:amount" sql:"type:decimal(32,16);"`
	Blockhash     string    `json:"blockhash"`
	Confirmations int32     `json:"confirmations"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	WalletID      int64     `json:"wallet_id"`
	CompanyID     int64     `json:"company_id"`
	Txid          string    `json:"txid"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
}

type DepositPagesQuery struct {
	PerPage int             `json:"per_page"`
	Page    int             `json:"page"`
	Query   *DepositQuery   `json:"query"`
	OrderBy *DepositOrderBy `json:"order_by"`
}

type DepositQuery struct {
	To          int64  `json:"to"`
	From        int64  `json:"from"`
	UpdatedTo   int64  `json:"updated_to"`
	UpdatedFrom int64  `json:"updated_from"`
	Currency    string `json:"currency"`
	CompanyID   int64  `json:"-"`
	WalletID    int64  `json:"wallet_id"`
}
type DepositOrderBy struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type DepositPagesResponse struct {
	Data  []*DepositItem `json:"data"`
	Page  int            `json:"page"`
	Pages int            `json:"pages"`
	Total int            `json:"total"`
}

// Deposit is a structure manager all about Deposit
type Deposit struct {
	client *Xchange
}

//Deposit - Instance de Deposit
func (c *Xchange) Deposit() *Deposit {
	return &Deposit{client: c}
}

func (x *Deposit) List(req DepositPagesQuery) (*DepositPagesResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *DepositPagesResponse
	err, errAPI := x.client.Request("POST", "/api/deposits", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
