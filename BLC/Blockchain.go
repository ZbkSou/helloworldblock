package BLC

type Blockchain struct {
	Blocks []*Block //储存区块
}

//创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	genesisBlock := CreateGenesisBlock("Genesis data")
	return &Blockchain{[]*Block{genesisBlock}}
}

//增加区块到链
func (blc *Blockchain) AddBlockToBlockchain(data string, height int64, preHash []byte) {
	//创建新区块
	newBlock := NewBlock(data, height, preHash)
	//增加到blc上
	blc.Blocks = append(blc.Blocks, newBlock)
}
