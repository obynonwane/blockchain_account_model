package main

import (
	"fmt"
	"log"

	"github.com/obynonwane/blockchain_account_model/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
	t := wallet.NewTrandaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Println("signature %s \n", t.GenerateSignature())
}
