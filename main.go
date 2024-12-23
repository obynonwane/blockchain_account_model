package main

import (
	"fmt"
	"log"

	"github.com/obynonwane/blockchain_account_model/block"
)

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	myBlockchainAddress := "my_blockchain_address"

	// main function initialise and create the first block
	// with a nonce of 0 and initail hash
	blockChain := block.NewBlockchain(myBlockchainAddress) // init block

	// subsequent blocks
	blockChain.AddTransaction("A", "B", 1.2)
	blockChain.Mining()

	blockChain.AddTransaction("C", "D", 5.2)
	blockChain.AddTransaction("E", "G", 10.2)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))

}
