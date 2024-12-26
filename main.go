package main

import (
	"fmt"
	"log"

	"github.com/obynonwane/blockchain_account_model/block"
	"github.com/obynonwane/blockchain_account_model/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	// walletM := wallet.NewWallet()
	// walletA := wallet.NewWallet()
	// walletB := wallet.NewWallet()

	// t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	// fmt.Printf("signature %s \n", t.GenerateSignature())

	// //Blockchain Node side
	// blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	// isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())
	// fmt.Println("Added?:", isAdded)

	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)
}
