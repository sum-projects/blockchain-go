package main

func main() {
	bc := NewBlockchain()
	bc.Print()

	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(5, previousHash)
	bc.Print()

	previousHash = bc.LastBlock().Hash()
	bc.CreateBlock(2, previousHash)
	bc.Print()
}
