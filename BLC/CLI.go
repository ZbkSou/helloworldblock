package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func printUsage() {

	fmt.Println("\nUsage:")
	fmt.Println("\t createblockchain -address -- 交易数据 ")
	fmt.Println("\t addblock -data DATA -- 交易数据")
	fmt.Println("\t printchain -- 输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

}

func (cli *CLI) addBlock(txs []*Transaction) {
	if DBExists() == false {
		fmt.Println("数据不存在....")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlockchain(txs)
}
func (cli *CLI) printchain() {
	if DBExists() == false {
		fmt.Println("数据不存在....")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.PrintChain()
}
func (cli *CLI) createGenesisBlockchain(address string) {
	CreateBlockchainWithGenesisBlock(address)
	BlockchainObject().DB.Close()
}
func (cli *CLI) Run() {
	isValidArgs()
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createblockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "zbk", "交易数据")

	flagCreateBlockchainWithAddress := createblockchainCmd.String("address", "Genesis block data ......", "创建创世块的地址")

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
	case "createblockchain":
		err := createblockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}
	//增添区块信息
	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagAddBlockData)
		cli.addBlock([]*Transaction{})
	}

	// 打印所有区块
	if printChainCmd.Parsed() {
		fmt.Println("\nprint all data")
		cli.printchain()
	}

	//创建区块链
	if createblockchainCmd.Parsed() {
		if *flagCreateBlockchainWithAddress == "" {
			fmt.Println("地址不能为空")
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagCreateBlockchainWithAddress)
		cli.createGenesisBlockchain(*flagCreateBlockchainWithAddress)
	}

}
