package BLC

type TXintput struct {
	//交易的id
	TxHash []byte
	//储存txoutput在vout的索引
	Vout int
	//用户签名
	ScriptSig string
}
