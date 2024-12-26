package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// accept command line argument
	port := flag.Uint("port", 5001, "TCP Port Number for Blockchain Server")
	// processes command line argument
	flag.Parse()

	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
