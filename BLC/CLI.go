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
	fmt.Println("\t send -from From -to to -amount amount -- 交易数据")
	fmt.Println("\t printchain -- 输出区块信息")
	fmt.Println("\t getbalance -- 查询账号余额")
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

//创建创世区块
func (cli *CLI) createGenesisBlockchain(address string) {
	blockchain := CreateBlockchainWithGenesisBlock(address)
	defer blockchain.DB.Close()
}

//发送功能
func (cli *CLI) send(from []string, to []string, amount []string) {
	if DBExists() == false {
		fmt.Println("数据不存在....")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)
}

func (cli *CLI) getBalance(address string) {
	UnSpentTransationsWithAdress(address)
}
func (cli *CLI) Run() {
	isValidArgs()
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createblockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getbalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	flagSendFrom := sendBlockCmd.String("from", "", "转账来源地址")
	flagSendTo := sendBlockCmd.String("to", "", "转账目的地址")
	flagSendAmount := sendBlockCmd.String("amount", "", "转账金额")

	flagCreateBlockchainWithAddress := createblockchainCmd.String("address", "Genesis block data ......", "创建创世块的地址")
	getbalanceCmdWithAdress := getbalanceCmd.String("address", "", "查询账号余额")
	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
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
	case "getbalance":
		err := getbalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}
	//处理转账信息
	if sendBlockCmd.Parsed() {
		if *flagSendFrom == "" || *flagSendTo == "" || *flagSendAmount == "" {
			printUsage()
			os.Exit(1)
		}

		fmt.Println()
		from := JSONToArray(*flagSendFrom)
		to := JSONToArray(*flagSendTo)
		amount := JSONToArray(*flagSendAmount)
		cli.send(from, to, amount)
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
		cli.createGenesisBlockchain(*flagCreateBlockchainWithAddress)
	}
	//创建区块链
	if getbalanceCmd.Parsed() {
		if *getbalanceCmdWithAdress == "" {
			fmt.Println("地址不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*getbalanceCmdWithAdress)
	}
}
