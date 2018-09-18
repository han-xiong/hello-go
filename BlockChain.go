package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (this *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10))
	headers := bytes.Join([][]byte{this.PrevBlockHash, this.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	this.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{}
	block.Timestamp = time.Now().Unix()
	block.Data = []byte(data)
	block.PrevBlockHash = prevBlockHash
	block.Hash = []byte{}
	block.SetHash()
	return &block
}

type BlockChain struct {
	Blocks []*Block
}

func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("Genesis blcok")
	genesisBlock.PrevBlockHash = []byte{}
	return &genesisBlock
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (this *BlockChain) AddBlock(data string) {
	prevBlock := this.Blocks[len(this.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	this.Blocks = append(this.Blocks, newBlock)
}

func main() {
	bc := NewBlockChain()

	var cmd string

	for {
		fmt.Println("print '1' add new block to chains")
		fmt.Println("print '2' show ")
		fmt.Println("print 'other' exit")
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "1":
			input := make([]byte, 1024)
			fmt.Println("input your data:")
			os.Stdin.Read(input)
			bc.AddBlock(string(input))
		case "2":
			for i, block := range bc.Blocks {
				fmt.Println("====================")
				fmt.Println("num ", i, ": ")
				fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Println("====================")
			}
		default:
			fmt.Println("exit")
		}
	}
}
