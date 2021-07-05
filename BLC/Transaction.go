package BLC

type Transaction struct {
	//1 交易hash
	TxHash []byte
	//2输入
	Vins []*TXintput
	//3 输出
	Vouts []*TXOutput
}
