package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)


type BlockOperation interface {
	GenerateBlock() Block
	CalculateHash() string
	IsBlockValid() bool
}


//GenerateBlock new block
func (b Block) GenerateBlock(BPM string) Block {

	var newBlock Block
	t := time.Now()

	newBlock.Index = b.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = b.Hash
	newBlock.Hash = newBlock.CalculateHash()
	return newBlock
}

// SHA256 hashing
func (b Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.BPM + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}


// IsBlockValid make sure block is valid by checking index, and comparing the hash of the previous block
func (b Block)IsBlockValid(oldBlock Block) bool {

	if oldBlock.Index+1 != b.Index {
		return false
	}

	if oldBlock.Hash != b.PrevHash {
		return false
	}

	if b.CalculateHash() != b.Hash {
		return false
	}

	return true
}
