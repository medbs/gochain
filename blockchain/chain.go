package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	index        int
	timeStamp    string
	hash         string
	previousHash string
}

func generateNewBlock(oldBlock Block) Block {
	var newBlock Block
	t := time.Now()
	newBlock.index = oldBlock.index + 1
	newBlock.timeStamp = t.String()
	newBlock.previousHash = oldBlock.hash
	return newBlock
}

// Calculate hash for a new block
func calculateHash(block Block) string {
	data := strconv.Itoa(block.index) + block.timeStamp + block.hash + block.previousHash
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isBlockValid(oldBlock Block, newBlock Block) bool {
	if (oldBlock.index != newBlock.index+1) || (newBlock.previousHash != oldBlock.hash) {
		return false
	} else if calculateHash(newBlock) != newBlock.hash {
		return false
	}
	return true
}
