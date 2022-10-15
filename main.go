package main

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
}
