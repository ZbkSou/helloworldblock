package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Blockchain *Blockchain
}

func printUsage() {

	fmt.Println("\nUsage:")
	fmt.Println("\t addblock -data DATA -- 交易数据")
	fmt.Println("\t printchain -- 输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

}

func (cli *CLI) addBlock(data string) {
	cli.Blockchain.AddBlockToBlockchain(data)

}
func (cli *CLI) printchain() {
	cli.Blockchain.PrintChain()
}
func (cli *CLI) Run() {
	isValidArgs()
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "zbk", "交易数据")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}
	if printChainCmd.Parsed() {
		fmt.Println("\nprint all data")
		cli.printchain()
	}

}
