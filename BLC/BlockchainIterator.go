package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//创建迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

func (blockchainIterator *BlockchainIterator) Next() *Block {
	var blcok *Block
	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockByters := b.Get(blockchainIterator.CurrentHash)
			//获取当前区块
			blcok = DeserializeBlock(currentBlockByters)
			blockchainIterator.CurrentHash = blcok.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return blcok
}
