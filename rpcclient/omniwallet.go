package rpcclient

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/qizikd/btcd/btcjson"
	"strconv"
)

type FutureGetOmniBalanceResult chan *response

type Omni_GetbalanceResult struct {
	Balance  string `json:"balance"`
	Reserved string `json:"reserved"`
	Frozen   string `json:"frozen"`
}

type Omni_ListtransactionResult struct {
	Txid             string `json:"txid"`
	Fee              string `json:"fee"`
	Sendingaddress   string `json:"sendingaddress"`
	Referenceaddress string `json:"referenceaddress"`
	Ismine           bool   `json:"ismine"`
	Version          int    `json:"version"`
	Type_int         int    `json:"type_int"`
	Type             string `json:"type"`
	Propertyid       int    `json:"propertyid"`
	Divisible        bool   `json:"divisible"`
	Amount           string `json:"amount"`
	Valid            bool   `json:"valid"`
	Blockhash        string `json:"blockhash"`
	Blocktime        int64  `json:"blocktime"`
	Positioninblock  int64  `json:"positioninblock"`
	Block            int64  `json:"block"`
	Confirmations    int64  `json:"confirmations"`
}

func (r FutureGetOmniBalanceResult) Receive() (int, error) {
	res, err := receiveFuture(r)
	if err != nil {
		fmt.Println("GetOmniBalanceResultError:",err)
		return 0, err
	}
	fmt.Println("GetOmniBalanceResult: ", string(res))
	// Unmarshal result as a floating point number.
	var omni_getbalanceResult Omni_GetbalanceResult
	err = json.Unmarshal(res, &omni_getbalanceResult)
	if err != nil {
		return 0, err
	}

	balance, err := strconv.ParseFloat(omni_getbalanceResult.Balance, 64)
	if err != nil {
		balance = 0
	}
	balance = balance * btcutil.SatoshiPerBitcoin
	if balance < 0 {
		return int(balance - 0.5), nil
	}
	return int(balance + 0.5), nil
}

type FutureOmniSendResult chan *response

func (r FutureOmniSendResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}
	fmt.Println("OmniSendResult: ", string(res))
	// Unmarshal result as a floating point number.
	var txhash string
	err = json.Unmarshal(res, &txhash)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

type FutureOmni_ListtransactionResult chan *response

func (r FutureOmni_ListtransactionResult) Receive() (result []Omni_ListtransactionResult, err error) {
	res, err := receiveFuture(r)
	if err != nil {
		return
	}
	fmt.Println("Omni_ListtransactionResult: ", string(res))
	// Unmarshal result as a floating point number.
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}
	return result, nil
}

func (c *Client) GetOminBalanceAsync(account string, propertyid int) FutureGetOmniBalanceResult {
	cmd := btcjson.NewOmni_GetbalanceCmd(&account, &propertyid)
	return c.sendCmd(cmd)
}

func (c *Client) GetOmniBalance(account string, propertyid int) (int, error) {
	return c.GetOminBalanceAsync(account, propertyid).Receive()
}

func (c *Client) Omni_ListtransactionsAsync(account string, count int, skip int) FutureOmni_ListtransactionResult {
	cmd := btcjson.NewOmni_ListtransactionsCmd(&account, &count, &skip)
	return c.sendCmd(cmd)
}

func (c *Client) Omni_Listtransactions(account string, count int, skip int) ([]Omni_ListtransactionResult, error) {
	return c.Omni_ListtransactionsAsync(account, count, skip).Receive()
}

func (c *Client) OminSendAsync(account string, toaccount string, amount string, propertyid int) FutureOmniSendResult {
	cmd := btcjson.NewOmni_SendCmd(&account, &toaccount, &propertyid, &amount)
	return c.sendCmd(cmd)
}

func (c *Client) OmniSend(account string, toaccount string, amount string, propertyid int) (string, error) {
	return c.OminSendAsync(account, toaccount, amount, propertyid).Receive()
}
