package response

import (
	"yam-api/source/utils/etherscan/helper"
	"encoding/json"
)

// @dev Carrier for every etherscan response.
type Envelope struct {
	Status int `json:"status,string"`
	Message string `json:"message"`
	Result json.RawMessage `json:"result"`
}

type Tx struct {
	BlockNumber       int     `json:"blockNumber,string"`
	TimeStamp         helper.Time    `json:"timeStamp"`
	Hash              string  `json:"hash"`
	Nonce             int     `json:"nonce,string"`
	BlockHash         string  `json:"blockHash"`
	TransactionIndex  int     `json:"transactionIndex,string"`
	From              string  `json:"from"`
	To                string  `json:"to"`
	Value             *helper.BigInt `json:"value"`
	Gas               int     `json:"gas,string"`
	GasPrice          *helper.BigInt `json:"gasPrice"`
	IsError           int     `json:"isError,string"`
	TxReceiptStatus   string  `json:"txreceipt_status"`
	Input             string  `json:"input"`
	ContractAddress   string  `json:"contractAddress"`
	CumulativeGasUsed int     `json:"cumulativeGasUsed,string"`
	GasUsed           int     `json:"gasUsed,string"`
	Confirmations     int     `json:"confirmations,string"`
	TokenName         string  `json:"tokenName"`
	TokenSymbol       string  `json:"tokenSymbol"`
	TokenDecimal      uint8   `json:"tokenDecimal,string"`
}
