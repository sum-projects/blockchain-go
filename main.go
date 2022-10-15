package main

import "fmt"

func main() {
	myAddress := "my_blockchain_address"
	bc := NewBlockchain(myAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	bc.Print()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	bc.Mining()
	bc.Print()

	fmt.Printf("my_blockchain_address %.1f\n", bc.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", bc.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", bc.CalculateTotalAmount("D"))
}
