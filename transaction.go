package main

import "encoding/json"

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{
		senderAddress:    sender,
		recipientAddress: recipient,
		value:            value,
	}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	m, _ := json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		Value            float32 `json:"value"`
	}{
		SenderAddress:    t.senderAddress,
		RecipientAddress: t.recipientAddress,
		Value:            t.value,
	})
	return m, nil
}
