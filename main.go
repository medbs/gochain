package main

import ledger "gochain/bc"

func main() {
	var chain ledger.Chain
	ledger.Launch(&chain)

}
