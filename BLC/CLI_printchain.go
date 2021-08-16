package BLC

import "fmt"

func printUsage() {

	fmt.Println("\nUsage:")
	fmt.Println("\t createblockchain -address -- 交易数据 ")
	fmt.Println("\t send -from From -to to -amount amount -- 交易数据")
	fmt.Println("\t printchain -- 输出区块信息")
	fmt.Println("\t getbalance -- 查询账号余额")
}
