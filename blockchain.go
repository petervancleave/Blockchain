package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)


type Block struct {
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}


type Blockchain struct {
	Blocks []*Block
}


func (b *Block) calculateHash() {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d%s%s%s%d", b.Timestamp, b.Data, b.PrevHash, b.Hash, b.Nonce)))
	calculatedHash := hash.Sum(nil)
	b.Hash = hex.EncodeToString(calculatedHash)
}


func (bc *Blockchain) createBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Hash:      "",
		Nonce:     0,
	}
	newBlock.mineBlock()
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}


func (b *Block) mineBlock() {
	for {
		b.Nonce++
		b.calculateHash()

		// consider it mined if the hash starts with "0000"
		if b.Hash[:4] == "0000" {
			fmt.Println("Block mined:", b.Hash)
			break
		}
	}
}

// create genesis block in the blockchain
func createGenesisBlock() *Block {
	return &Block{
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
		Nonce:     0,
	}
}

func createBlockchain() *Blockchain {
	genesisBlock := createGenesisBlock()
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

func main() {

	blockchain := createBlockchain()

	// add blocks
	blockchain.createBlock("Block 1 Data")
	blockchain.createBlock("Block 2 Data")
	blockchain.createBlock("Block 3 Data")

	// print blocks
	for _, block := range blockchain.Blocks {
		fmt.Printf("Timestamp: %d, Data: %s, PrevHash: %s, Hash: %s, Nonce: %d\n",
			block.Timestamp, block.Data, block.PrevHash, block.Hash, block.Nonce)
	}
}
