package main

import (
	"fmt"
	"assignment01bca/assignment01bca"
)

func main() {
	var blockchain []*assignment01bca.Block

	// Requirement: Genesis Block with Transaction = 'Genesis Block - 1574'
	// Requirement: Nonce = Sum of digits (1+5+7+4 = 17)
	genesis := assignment01bca.NewBlock("Genesis Block - 1574", 17, "0", 0)
	blockchain = append(blockchain, genesis)

	// Requirement: Create at least 3 blocks total
	// Requirement: Roll number digits (574) included in at least one block transaction
	block1 := assignment01bca.NewBlock("Alice to Bob", 101, genesis.Hash, 1)
	blockchain = append(blockchain, block1)

	block2 := assignment01bca.NewBlock("Transfer to student 574", 202, block1.Hash, 2)
	blockchain = append(blockchain, block2)

	// 1. Print the initial blockchain
	fmt.Println("=== INITIAL BLOCKCHAIN ===")
	assignment01bca.ListBlocks(blockchain)

	// 2. Verify the initial blockchain
	isValid := assignment01bca.VerifyChain(blockchain)
	fmt.Printf("Is Chain Valid? %v\n\n", isValid)

	// 3. Tamper with Block 1 using ChangeBlock()
	fmt.Println("=== TAMPERING WITH BLOCK 1 ===")
	assignment01bca.ChangeBlock(block1, "Hacker changed the transaction!")
	
	// 4. Run VerifyChain() again to show failure
	isValidAfterTamper := assignment01bca.VerifyChain(blockchain)
	fmt.Printf("Is Chain Valid after tampering? %v\n", isValidAfterTamper)
}