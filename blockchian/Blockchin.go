package blockchian

type BlockChain struct {
	Blocks []*Block // stored sort block
}

// 1. create block with genesis block
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("genesis block .....")
	return &BlockChain{[]*Block{genesisBlock}}
}

// 2 add block
func (blc *BlockChain) AddBlockToBlockChain(data string,height int64,preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	blc.Blocks = append(blc.Blocks, newBlock)
}
