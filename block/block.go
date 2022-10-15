package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/sum-project/blockchain/transaction"
	"time"
)

type Block struct {
	Nonce        int
	PreviousHash [32]byte
	Timestamp    int64
	Transactions []*transaction.Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transaction []*transaction.Transaction) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transaction,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64                      `json:"timestamp"`
		Nonce        int                        `json:"Nonce"`
		PreviousHash [32]byte                   `json:"previous_hash"`
		Transaction  []*transaction.Transaction `json:"transaction"`
	}{
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
		Transaction:  b.Transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.Timestamp)
	fmt.Printf("Nonce                 %d\n", b.Nonce)
	fmt.Printf("previous_hash         %x\n", b.PreviousHash)
	for _, t := range b.Transactions {
		t.Print()
	}
}
