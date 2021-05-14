package main

import "fmt"
import "helloworldblock/BLC"

func main() {

	Blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//创世区块
	fmt.Println(Blockchain.Blocks)
	//	新区快
	Blockchain.AddBlockToBlockchain("send 100RMB to zhangsan", Blockchain.Blocks[len(Blockchain.Blocks)-1].Height+1, Blockchain.Blocks[len(Blockchain.Blocks)-1].Hash)
	Blockchain.AddBlockToBlockchain("send 300RMB to zhangsan", Blockchain.Blocks[len(Blockchain.Blocks)-1].Height+1, Blockchain.Blocks[len(Blockchain.Blocks)-1].Hash)
	Blockchain.AddBlockToBlockchain("send 500RMB to zhangsan", Blockchain.Blocks[len(Blockchain.Blocks)-1].Height+1, Blockchain.Blocks[len(Blockchain.Blocks)-1].Hash)
	fmt.Println(Blockchain.Blocks)
	bytes := Blockchain.Blocks[1].Serialize()
	fmt.Println(bytes)
	block := BLC.DeserializeBlock(bytes)
	fmt.Print(block)

}
