package bc

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

//GenerateNewBlock new block
func GenerateNewBlock(oldBlock Block) Block {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	return newBlock
}

//GenerateBlock new block
func GenerateBlock(oldBlock Block, BPM int) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

// CalculateHash Calculate SHA256 hash for a new block
/*func CalculateHash(block Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp + block.Hash + block.PrevHash
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}*/

// SHA256 hashing
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}



// IsBlockValid make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock , oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
