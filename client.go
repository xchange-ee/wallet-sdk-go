package xchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Xchange struct {
	client *http.Client
	Token  string
	Env    string
}

type Error struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}

func NewClient(Token, env string) *Xchange {
	return &Xchange{
		client: &http.Client{Timeout: 60 * time.Second},
		Token:  Token,
		Env:    env,
	}
}

func (x *Xchange) Request(method, action string, body []byte, out interface{}) (error, *Error) {
	if x.client == nil {
		x.client = &http.Client{Timeout: 60 * time.Second}
	}
	url := x.devProd()
	endpoint := fmt.Sprintf("%s%s", url, action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}

	log.Printf("endpoint %s\n", endpoint)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", x.Token))
	res, err := x.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return err, nil
		}
		errAPI.Body = string(bodyResponse)
		return nil, &errAPI
	}
	err = json.Unmarshal(bodyResponse, out)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

func (Xchange *Xchange) devProd() string {
	if Xchange.Env == "develop" {
		return "http://api.sandbox.sc.xchange.ee/node/v1"
	}
	return "https://api.sc.xchange.ee/node/v1"
}
