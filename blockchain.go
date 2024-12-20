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
func(b *Block) Print(){
	fmt.Printf("timestamp")
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	// create a new block
	b := NewBlock(0, "init hash")

	fmt.Println(b)
}
