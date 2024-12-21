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

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}

	fmt.Printf("%s\n", strings.Repeat("*", 25))
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
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
	// main function initialise and create the first block
	// with a nonce of 0 and initail hash
	// blockChain := NewBlockchain() // init block

	// subsequent blocks
	// blockChain.CreateBlock(5, "hash 1")
	// blockChain.CreateBlock(2, "hash 2")
	// blockChain.Print()

}
