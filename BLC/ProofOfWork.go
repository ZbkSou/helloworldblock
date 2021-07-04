package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//难度
const targetBit = 16

type ProofOfWork struct {
	Block  *Block //当前要验证的区块
	target *big.Int
}

//创建新的工作量证明
func NewProofOfWork(block *Block) *ProofOfWork {
	//1.创建1的target
	//2.左移256 - targetBit位
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}

//区块数据拼接到一个数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevBlockHash,
		pow.Block.HashTransactions(),
		IntToHex(pow.Block.Timestamp),
		IntToHex(int64(targetBit)),
		IntToHex(int64(nonce)),
		IntToHex(int64(pow.Block.Height)),
	}, []byte{})
	return data
}

//算出满足工作量证明的hsh
func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	//1.blockP拼接数组

	//2.生成hash
	//3.判断有效性
	nonce := 0
	var hashInt big.Int
	var hash [32]byte
	for true {
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		//通过比较大小来判断是否满足要求工作量证明
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce += 1
	}
	return hash[:], int64(nonce)
}

//判断当前hash,只与工作量正面来比较 是否有效
func (proofOfWork *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)
	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}
