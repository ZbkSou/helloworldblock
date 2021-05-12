package BLC

import (
	"fmt"
	"time"
)

type Block struct {
	//	区块高度
	Height int64
	//上一个区块的hash
	PrevBlockHash []byte
	//	交易数据
	Data []byte
	//	时间戳
	Timestamp int64
	//	hash
	Hash []byte
	//	Nonce
	Nonce int64
}

// NewBlock 创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, prevBlockHash,
		[]byte(data), time.Now().Unix(), nil, 0}

	//工作量证明
	pow := NewProofOfWork(block)
	fmt.Println()
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
