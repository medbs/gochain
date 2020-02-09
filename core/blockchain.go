package core

//P2pConfig represents the configuration of the BlockChain node
type P2pConfig struct {
	ListenF int
	Target string
	Secio bool
	Seed int64
}


//Block represents each 'item' in the BlockChain
type Block struct {
	Index     int
	Timestamp string
	BPM       string
	Hash      string
	PrevHash  string
}

//Chain represents the BlockChain
type Chain struct {
	BlockChain []Block
	P2pConfig *P2pConfig
}

//returns a configured BlockChain
func NewBlockChain(c *P2pConfig) *Chain {

	return &Chain {
		P2pConfig:c,
	}
}