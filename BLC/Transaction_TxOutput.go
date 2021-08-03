package BLC

type TXOutput struct {
	Value        int64
	ScriptPubKey string
}

//判断当前消费归属
func (txOutput *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {
	return txOutput.ScriptPubKey == address

}
