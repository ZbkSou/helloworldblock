package BLC

type TXintput struct {
	//来源订单交易的id
	TxHash []byte
	//储存txoutput在vout的索引
	Vout int
	//用户签名
	ScriptSig string
}

//判断当前消费归属
func (txInput *TXintput) UnLockWithAddress(address string) bool {
	return txInput.ScriptSig == address

}
