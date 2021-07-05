package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blockchain.db"
const blockTableName = "blockchain"

type Blockchain struct {
	//Blocks []*Block //储存区块，之后需要改成持久化存贮
	Tip []byte   //最新区块的hash值
	DB  *bolt.DB //数据库
}

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

//创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock(txs []*Transaction) *Blockchain {
	//判断数据库是否存在
	if DBExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1)
	}
	//创建创世区块
	fmt.Println("创建创世区块...")
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		//检查表是否存在
		if b == nil {
			//创建表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
			}
		}
		if b != nil {
			// 创建创世区块
			genesisBlock := CreateGenesisBlock(txs)
			// 将创世区块储存到表
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//	存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

		}
		return nil
	})
	return nil
}

//增加区块到链
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取数据库最新区块
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)
			//创建区块
			newBlock := NewBlock(txs, block.Height+1, block.Hash)
			//保存到数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//更新tip 数据库+内存
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

//遍历打印所有区块信息
func (blc *Blockchain) PrintChain() {
	blockchainIterator := blc.Iterator()
	for {
		block := blockchainIterator.Next()
		fmt.Println("===================")
		fmt.Printf("height : %d\n", block.Height)
		fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %v\n", block.Txs)
		fmt.Printf("Timestamp : %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Println("===================")
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

//得到区块链
func BlockchainObject() *Blockchain {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	return &Blockchain{tip, db}

}
