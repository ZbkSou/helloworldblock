package main

import (
	"flag"
	"fmt"
)

func main() {
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	flag.Parse()
	fmt.Printf("%s\n", *flagPrintChainCmd)
}
