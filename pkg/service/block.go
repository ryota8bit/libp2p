package service

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Point     int
	Hash      string
	PrevHash  string
}

// BlockChain is a series of validated Blocks
var BlockChain []Block

// IsBlockValid make sure block is valid by checking index,
// comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
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

// SHA256 hashing
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) +
		block.Timestamp +
		strconv.Itoa(block.Point) +
		block.PrevHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func GenerateBlock(oldBlock Block, Point int) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Point = Point
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}
