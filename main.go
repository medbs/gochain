package main

import (
	ledger "gochain/core"
)

func main() {
	var chain ledger.Chain
	ledger.Launch(&chain)

}
