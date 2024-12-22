package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"time"
)

// 1. sha256 - Hash algorithm

// block struct
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

// create contructor function to create a new block
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash

	return b
}

// helper function to print block
func (b *Block) Print() {
	fmt.Printf("timestamp            %d\n", b.timestamp)
	fmt.Printf("nonce                %d\n", b.nonce)
	fmt.Printf("previous_hash        %x\n", b.previousHash)
	fmt.Printf("transactions         %s\n", b.transactions)
}

// creating the blochain
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc

}

// create a blockchain by
// a. creating a new block, using the nonce and previoushash
// b. appending the new block to the chain
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

// return the last block in the chain
func (bc *Blockchain) LastBlock() *Block {
	// go into the array and pick the last block
	// e.g bc.chain[10]
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}

	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

// struct for transaction
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// function for creating new transaction
func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// print transaction
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" senderBlockchainAddress           %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipientBlockchainAddress        %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                             %.1f\n", t.value)
}

// json marshal transaction for hashing
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.recipientBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func (b *Block) Hash() [32]byte {
	// make the block into a json formatted string
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	// return the sha256 byte as the hash of the block
	return sha256.Sum256([]byte(m))
}

// special function in GOLANG for creating custom encoding/json
// which is invoked by json.Marshal() function
// part of the json.Marshaler interface
// which allows you to define a custom way to convert your struct into JSON
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {

	// main function initialise and create the first block
	// with a nonce of 0 and initail hash
	blockChain := NewBlockchain() // init block

	// subsequent blocks
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(1, previousHash)

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()

}
