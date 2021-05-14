package main

import (
	"helloworldblock/BLC"
)

func main() {

	Blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer Blockchain.DB.Close()

	////	新区快
	Blockchain.AddBlockToBlockchain("send 100RMB to zhangsan")
	Blockchain.AddBlockToBlockchain("send 300RMB to zhangsan")
	Blockchain.AddBlockToBlockchain("send 500RMB to zhangsan")
	Blockchain.PrintChain()

}
