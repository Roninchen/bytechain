package blockchian

import (
	"fmt"
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块HASH
	PrevBlockHash []byte
	//3. 交易数据
	Data []byte
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
	// Nonce
	Nonce int64
}

//1. new block
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {

	//block
	block := &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil,0}

	// pow return hash nonce
	pow :=NewProofOfWork(block)

	//000000
	hash,nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	fmt.Println()

	return block

}

// 2. create genesis block
func CreateGenesisBlock(data string) *Block{
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
