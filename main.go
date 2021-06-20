package main

import (
	"./BLC"

)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
	cli := BLC.CLI{blockchain}
	cli.Run()
}
