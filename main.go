package main

import (
	"blockchain/chain"
	"time"
)

var blockchain []Block

func main() {

	t := time.Now()
	genesisBlock := chain.Block{}
	genesisBlock = Block{0, t.String(), chain.calculatHash(genesisBlock), ""}
	blockchain = append(blockchain, genesisBlock)

}
