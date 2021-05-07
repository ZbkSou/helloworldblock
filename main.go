package main
import "fmt"
import "./BLC"
func main() {

	intVal :=BLC.CreateGenesisBlock("Genenis Block")

	fmt.Println(intVal)
}
