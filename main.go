package main

func main() {
	bc := NewBlockchain()
	bc.Print()
	bc.CreateBlock(5, "hash 1")
	bc.Print()
	bc.CreateBlock(2, "hah 2")
	bc.Print()
}
