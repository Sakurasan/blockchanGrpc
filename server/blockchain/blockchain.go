package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

//块
type Block struct {
	Data          string
	PrevBlockHash string
	Hash          string
}

func (t *Block) CalculateHash() {
	blockdata := t.PrevBlockHash + t.Data
	hash := sha256.Sum256([]byte(blockdata))
	t.Hash = hex.EncodeToString(hash[:])
}

func GenerateNewBlock(preblock Block, data string) Block {
	newBlock := Block{}
	newBlock.Data = data
	newBlock.PrevBlockHash = preblock.Hash
	newBlock.CalculateHash()
	return newBlock
}

func GenerateBlock() Block {
	block := Block{}
	block.Data = "创世块"
	block.CalculateHash()
	return block
}

//链
type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	generateBlock := GenerateBlock()
	newblockchain := new(BlockChain)

	newblockchain.Blocks = append(newblockchain.Blocks, &generateBlock)
	return newblockchain
}

func (t *BlockChain) AddBlock(data string) {
	block := GenerateNewBlock(*t.Blocks[len(t.Blocks)-1], data)
	t.Blocks = append(t.Blocks, &block)
}
