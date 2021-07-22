package xchange

import (
	"encoding/json"
	"fmt"
	"time"
)

type AddressResponse struct {
	ID              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Address         string    `json:"address"`
	Currency        string    `json:"currency"`
	CompanyID       int64     `gorm:"column:company_id" json:"-"`
	Source          string    `json:"-"`
	WalletCustodyID int64     `json:"-"`
	WalletID        int64     `json:"wallet_id"`
}

type AddressRequest struct {
	WalletID int64 `json:"wallet_id"`
}

type WalletResponse struct {
	ID        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Currency  string    `json:"currency"`
	Balance   string    `json:"balance"`
}

type WalletRequest struct {
	Currency  string `json:"currency"`
	CompanyID int64  `gorm:"column:company_id" json:"-"`
}

// Wallet is a structure manager all about Wallet
type Wallet struct {
	client *Xchange
}

//Wallet - Instance de Wallet
func (c *Xchange) Wallet() *Wallet {
	return &Wallet{client: c}
}

//Create - create a new wallet
func (x *Wallet) Create(req WalletRequest) (*WalletResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *WalletResponse
	err, errAPI := x.client.Request("POST", "/api/wallet", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//GetWallet - create a new wallet
func (x *Wallet) Get(id int32) (*WalletResponse, *Error, error) {
	var response *WalletResponse
	err, errAPI := x.client.Request("GET", fmt.Sprintf("/api/wallet/%d", id), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//GetNewAddress - get new address
func (x *Wallet) GetNewAddress(req AddressRequest) (*AddressResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *AddressResponse
	err, errAPI := x.client.Request("POST", "/api/address", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
