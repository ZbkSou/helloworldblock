package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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
	txOutput := &TXOutput{1000, address}
	txCoinBase := &Transaction{
		[]byte{},
		[]*TXintput{txInput},
		[]*TXOutput{txOutput},
	}
	//设置hash值
	txCoinBase.HashTransaction()
	return txCoinBase

}

//对 transaction 序列化之后hash
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

//1.2 转账产生transaction

func NewSimpleTransaction(from string, to string, amount int) *Transaction {

	//找到from下所有未花费交易
	unSpentTx := UnSpentTransationsWithAdress(from)

	var txInputs []*TXintput
	var txOutputs []*TXOutput
	//来源
	b, _ := hex.DecodeString("")
	txInput := &TXintput{b, 10, from}

	txInputs = append(txInputs, txInput)
	//消费
	txOutput := &TXOutput{4, to}
	txOutputs = append(txOutputs, txOutput)
	//找零
	txOutput = &TXOutput{10 - 4, from}
	txOutputs = append(txOutputs, txOutput)
	tx := &Transaction{
		[]byte{},
		txInputs,
		txOutputs,
	}
	//设置hash值
	tx.HashTransaction()
	return tx
}
