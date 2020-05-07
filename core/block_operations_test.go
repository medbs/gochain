package core

import (
	"testing"
)

func TestGenerateBlock(t *testing.T) {

	b1 := Block{
		Index:     1,
		Timestamp: "stamp",
		BPM:       "bpm",
		Hash:      "41064f80255aee51e0e954f42ec238099eb636fbf11ffb929ecdb809d5d0dc6f",
		PrevHash:  " ",
	}

	generatedBlock := GenerateBlock(b1, "bpm2")

	if generatedBlock.PrevHash != b1.Hash {
		t.Errorf("the generated block previous has is not equal to the current block hash")
	}

	if generatedBlock.Index != b1.Index+1 {
		t.Errorf("the generated block index should be previous index + 1 ")
	}

	if generatedBlock.BPM != "bpm2" {
		t.Errorf("BPM value not assigned correctly")
	}
}

func TestCalculateHash(t *testing.T) {
	b1 := Block{
		Index:     1,
		Timestamp: "stamp",
		BPM:       "bpm",
		Hash:      "hash",
		PrevHash:  "prevhash",
	}
	//the fixed and and correct value of block hash
	correctHash := "41064f80255aee51e0e954f42ec238099eb636fbf11ffb929ecdb809d5d0dc6f"
	blockHash := CalculateHash(b1)

	if blockHash != correctHash {
		t.Errorf("returned value is not correct")
	}

}

func TestIsBlockValid(t *testing.T) {
	b1 := Block{
		Index:     1,
		Timestamp: "stamp",
		BPM:       "bpm",
		Hash:      "41064f80255aee51e0e954f42ec238099eb636fbf11ffb929ecdb809d5d0dc6f",
		PrevHash:  "prevhash",
	}

	b2 := Block{
		Index:     2,
		Timestamp: "stamp",
		BPM:       "bpm",
		Hash:      "4813e8a674e20be91754a8a92b8e01f35767fbdd13dc988fe0260943d1f17e1a",
		PrevHash:  "41064f80255aee51e0e954f42ec238099eb636fbf11ffb929ecdb809d5d0dc6f",
	}

	if !IsBlockValid(b2, b1) {
		t.Errorf("next block cannot be validated")
	}
}
