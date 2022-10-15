package main

import (
	"fmt"
	"strings"
)

func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previous_hash         %s\n", b.previousHash)
	fmt.Printf("transactions          %s\n", b.transactions)
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}
