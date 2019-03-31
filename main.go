package main

import (
	"bytechain/blockchian"
)

func main() {
	blockchian := blockchian.CreateBlockChainWithGenesisBlock()

	//new block

	blockchian.AddBlockToBlockChain("send 100 RMB to TOM", blockchian.Blocks[len(blockchian.Blocks)-1].Height+1, blockchian.Blocks[len(blockchian.Blocks)-1].Hash)

	blockchian.AddBlockToBlockChain("send 100 RMB to Bob", blockchian.Blocks[len(blockchian.Blocks)-1].Height+1, blockchian.Blocks[len(blockchian.Blocks)-1].Hash)

	blockchian.AddBlockToBlockChain("send 100 RMB to Alia", blockchian.Blocks[len(blockchian.Blocks)-1].Height+1, blockchian.Blocks[len(blockchian.Blocks)-1].Hash)

	blockchian.AddBlockToBlockChain("send 100 RMB to Ada", blockchian.Blocks[len(blockchian.Blocks)-1].Height+1, blockchian.Blocks[len(blockchian.Blocks)-1].Hash)
}