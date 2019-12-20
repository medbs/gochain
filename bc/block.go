package bc

//Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

//Chain represents the ledger
type Chain struct {
	Blockchain []Block
}
