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
func CreateBlockchainWithGenesisBlock() *Blockchain {
	if dbExists() {
		fmt.Println("创世区块已经存在")
		db, err := bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			hash := b.Get()
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
	}
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockHash []byte
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
			genesisBlock := CreateGenesisBlock("Genesis data")
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//	存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	return &Blockchain{blockHash, db}
}

//增加区块到链
func (blc *Blockchain) AddBlockToBlockchain(data string) {
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取数据库最新区块
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)
			//创建区块
			newBlock := NewBlock(data, block.Height+1, block.Hash)
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

func dbExists() bool {
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
		fmt.Printf("Data : %s\n", block.Data)
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
