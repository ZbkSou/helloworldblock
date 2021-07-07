package BLC

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
	}
	return &Transaction{}

}

//1.2 转账产生transaction
