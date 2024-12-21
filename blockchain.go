package main

import (
	"fmt"
	"log"
	"time"
)

// block struct
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// create contructor function to create a new block
func NewBlock(nonce int, previousHash string) *Block {
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
	fmt.Printf("previous_hash        %s\n", b.previousHash)
	fmt.Printf("transactions         %s\n", b.transactions)
}

// creating the blochain
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// create a blockchain by
// a. creating a new block, using the nonce and previoushash
// b. appending the new block to the chain
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	// main function initialise and create the first block
	// with a nonce of 0 and initail hash
	b := NewBlock(0, "init hash")
	b.Print()

}
