package main

import (
	ledger "gochain/bc"
)

// Blockchain is a series of validated Blocks
//var Blockchain []ledger.Block
var chain ledger.Chain

func main() {
	ledger.Launch(chain)
}
