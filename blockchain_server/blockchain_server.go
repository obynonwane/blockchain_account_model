package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/obynonwane/blockchain_account_model/block"
	"github.com/obynonwane/blockchain_account_model/wallet"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port: port,
	}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	// check the cache for key "blockchain"
	bc, ok := cache["blockchain"]
	if !ok {
		// create new minner wallet
		minnersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minnersWallet.BlockchainAddress(), bcs.port)
		// added created blockchain to the cache
		cache["blockchain"] = bc

		log.Printf("private_key %v", minnersWallet.PrivateKeyStr())
		log.Printf("pblic_key %v", minnersWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minnersWallet.BlockchainAddress())
	}

	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:])) /*converts byte slice to string */
	default:
		log.Printf("ERROR: Invalid http method")
	}
}

func (bcs *BlockchainServer) Run() {
	// registers handler function for a given url or pattern
	http.HandleFunc("/", bcs.GetChain)

	// starting the server and also converting the port to string while appending it to the server address
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.port)), nil))
}
