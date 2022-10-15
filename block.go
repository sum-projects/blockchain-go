package main

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transaction  []string `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transaction:  b.transactions,
	})
}
