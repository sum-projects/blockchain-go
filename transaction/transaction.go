package transaction

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	SenderAddress    string
	RecipientAddress string
	Value            float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{
		SenderAddress:    sender,
		RecipientAddress: recipient,
		Value:            value,
	}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	m, _ := json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		Value            float32 `json:"Value"`
	}{
		SenderAddress:    t.SenderAddress,
		RecipientAddress: t.RecipientAddress,
		Value:            t.Value,
	})
	return m, nil
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_address      %s\n", t.SenderAddress)
	fmt.Printf(" recipient_address   %s\n", t.RecipientAddress)
	fmt.Printf(" Value               %.1f\n", t.Value)
}
