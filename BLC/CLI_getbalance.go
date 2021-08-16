package BLC

import "fmt"

func (cli *CLI) getBalance(address string) {
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	amount := blockchain.GetBalance(address)
	fmt.Printf("%s 一共有%d 个token", address, amount)
}
