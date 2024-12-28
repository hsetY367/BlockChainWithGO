package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevBlock []byte
	Hash      []byte
}

type Blockchain struct {
	blocks []*Block
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivy")
	bc.AddBlock("Send 2 BTC to Ivy")

	for _, block := range bc.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevBlock)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlock, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlock, []byte{}}
	block.SetHash()
	return block
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
