package main

import "fmt"

func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previous_hash         %s\n", b.previousHash)
	fmt.Printf("transactions          %s\n", b.transactions)
}
