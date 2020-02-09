package core

type P2pConfig struct {
	ListenF int
	Target string
	Secio bool
	Seed int64
}


//Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       string
	Hash      string
	PrevHash  string
}

//Chain represents the ledger
type Chain struct {
	BlockChain []Block
	P2pConfig *P2pConfig
}


func NewBlockChain(c *P2pConfig) *Chain {

	return &Chain {
		P2pConfig:c,
	}
}