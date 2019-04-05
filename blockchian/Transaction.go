package blockchian

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// UTXO
type Transaction struct {

	//1. 交易hash

	TxHash []byte

	//2. 输入
	Vins []*TXInput

	//3. 输出
	Vouts []*TXOutput
}

//1. Transaction 创建分两种情况
//1. 创世区块创建时的Transaction
func NewCoinbaseTransaction(address string) *Transaction {
	//代表消费
	txInput := &TXInput{[]byte{},-1,"Genesis Data"}


	txOutput := &TXOutput{10,address}
	txCoinbase := &Transaction{[]byte{},[]*TXInput{txInput},[]*TXOutput{txOutput}}

	txCoinbase.HashTransaction()

	return txCoinbase
}

func (tx *Transaction) HashTransaction()  {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic()
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}