package main

import "fmt"
import "helloworldblock/BLC"

func main() {

	intVal := BLC.CreateGenesisBlock("Genenis Block")
	fmt.Println(intVal)
}
