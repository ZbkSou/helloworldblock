package BLC

import (
	"bytes"
	"encoding/ascii85"
	"encoding/gob"
)

//UTXO
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
	txCoinBase.TxHash
	return txCoinBase

}

func (tx *Transaction) HashTransaction() []byte {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err := encode.Encode(tx)

}

func (tx *Transaction) Serialize() []byte {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err := encode.Encode()

}

//1.2 转账产生transaction
