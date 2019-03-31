package blockchian

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//256位hash需要16个零
const targetBit  =  16


type ProofOfWork struct {
	Block *Block
	target *big.Int //大数存储
}

// create new ProofOfWork obj
func NewProofOfWork(block *Block) *ProofOfWork {

	//1. big.Int对象 1
	target := big.NewInt(1)

	//2. 左移256 - targetBit
	target = target.Lsh(target ,256 - targetBit)

	return &ProofOfWork{block,target}
}

func (proofOfWork *ProofOfWork) Run()([]byte,int64) {
	//1. Bloch param to byte array
	//2. hash
	//3. judge hash
	nonce := 0
	var hashInt big.Int
	var hash [32]byte
	for  {
		//prepare
		dataBytes := proofOfWork.prepareData(nonce)
		hash := sha256.Sum256(dataBytes)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])

		//judge
		if proofOfWork.target.Cmp(&hashInt) == 1{
			break
		}
		nonce = nonce +1
	}
	return hash[:],int64(nonce)
}

func (proofOfWork *ProofOfWork) IsValid() bool{
	//1.proofOfWork.Block.Hash
	//2.proofOfWork.Target
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)
	if proofOfWork.target.Cmp(&hashInt) == 1{
		return true
	}
	return false
}
// 数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)

	return data
}



