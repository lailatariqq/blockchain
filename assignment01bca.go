package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
	Index        int
	Timestamp    int64
}

// NewBlock creates a new block and adds it to the chain
func NewBlock(transaction string, nonce int, previousHash string, index int) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Index:        index,
		Timestamp:    time.Now().Unix(),
	}
	block.Hash = CalculateHash(block)
	return block
}

// CalculateHash calculates hash using SHA-256 and concatenates: 
// transaction + nonce + previousHash + index + timestamp
func CalculateHash(b *Block) string {
	data := b.Transaction + strconv.Itoa(b.Nonce) + b.PreviousHash + strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// ListBlocks prints all the blocks in a nice format
func ListBlocks(chain []*Block) {
	for _, b := range chain {
		fmt.Printf("Index:         %d\n", b.Index)
		fmt.Printf("Transaction:   %s\n", b.Transaction)
		fmt.Printf("Nonce:         %d\n", b.Nonce)
		fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
		fmt.Printf("Current Hash:  %s\n", b.Hash)
		fmt.Printf("------------------------------------------------------------\n")
	}
}

// ChangeBlock changes the transaction of a given block reference
func ChangeBlock(b *Block, newTransaction string) {
	b.Transaction = newTransaction
}

// VerifyChain verifies the blockchain integrity
func VerifyChain(chain []*Block) bool {
	for i := 0; i < len(chain); i++ {
		// 1. Check if the block's own data matches its hash
		if chain[i].Hash != CalculateHash(chain[i]) {
			fmt.Printf("Verification Failed: Block %d data has been tampered with.\n", i)
			return false
		}
		// 2. Check if the link to the previous block is intact
		if i > 0 {
			if chain[i].PreviousHash != chain[i-1].Hash {
				fmt.Printf("Verification Failed: Block %d link to previous block is broken.\n", i)
				return false
			}
		}
	}
	return true
}