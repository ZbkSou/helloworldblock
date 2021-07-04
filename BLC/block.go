package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//	区块高度
	Height int64
	//上一个区块的hash
	PrevBlockHash []byte
	//	交易数据
	Tx []*Transaction
	//	时间戳
	Timestamp int64
	//	hash
	Hash []byte
	//	Nonce
	Nonce int64
}

func (block *Block) HashTransactions() []byte {
	var
	return nil
}

// NewBlock 创建新的区块
func NewBlock(txs []*Transaction, height int64, prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, prevBlockHash,
		txs, time.Now().Unix(), nil, 0}

	//工作量证明
	pow := NewProofOfWork(block)
	fmt.Println()
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//创建创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

//将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err := encode.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block

}
