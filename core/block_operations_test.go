package core

import "testing"

func TestGenerateBlock(t *testing.T) {

	b1 := Block{
		Index:     1,
		Timestamp: "stamp",
		BPM:       "bpm",
		Hash:      "hash",
		PrevHash:  "prevhash",
	}

	generatedBlock := GenerateBlock(b1, "bpm2")

	if generatedBlock.PrevHash != b1.Hash {
		t.Errorf("the generated block previous has is not equal to the current block hash")
	}

	if generatedBlock.Index != b1.Index+1 {
		t.Errorf("the generated block index should be previous index + 1 ")
	}

	if generatedBlock.BPM != "bmp2" {
		t.Errorf("BPM value not assigned correctly")
	}
}

func TestCalculateHash(t *testing.T){

}
