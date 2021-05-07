package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

// SetHash 设置区块的哈希

func (block *Block) SetHash() {
	//	height转化成byte
	heightBytes := IntToHex(block.Height)
	// 将时间戳转化成byte
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)

	//拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})

	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

// NewBlock 创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, prevBlockHash,
		[]byte(data), time.Now().Unix(), nil}
	//给区块生成hash
	block.SetHash()
	return block
}

func  CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
