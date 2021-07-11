package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//UTXO
//结构体大写公开访问
type Transaction struct {
	//1 交易hash
	TxHash []byte
	//2输入
	Vins []*TXintput
	//3 输出
	Vouts []*TXOutput
}

//1 Transaction 创建分两种情况
//1.1 区块链创建时transaction

func NewCoinBaseTransaction(address string) *Transaction {

	txInput := &TXintput{[]byte{}, -1, "Genesis Data"}
	txOutput := &TXOutput{10, address}
	txCoinBase := &Transaction{
		[]byte{},
		[]*TXintput{txInput},
		[]*TXOutput{txOutput},
	}
	//设置hash值
	txCoinBase.HashTransaction()
	return txCoinBase

}

//
func (tx *Transaction) HashTransaction() {
	//作为结果接收变量
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err := encode.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

func (tx *Transaction) Serialize() []byte {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err := encode.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()

}

//1.2 转账产生transaction
