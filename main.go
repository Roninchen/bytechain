package main

import (
	"fmt"
	"bytechain/blockchian"
)

func main() {
	blockchian := blockchian.CreateBlockChainWithGenesisBlock()

	fmt.Println(blockchian.Blocks[0])

	//new block

	blockchian.AddBlockToBlockChain("send 100 RMB to TOM",blockchian.Blocks[len(blockchian.Blocks)-1].Height+1,blockchian.Blocks[len(blockchian.Blocks)-1].Hash)

	blockchian.AddBlockToBlockChain("send 100 RMB to Bob",blockchian.Blocks[len(blockchian.Blocks)-1].Height+1,blockchian.Blocks[len(blockchian.Blocks)-1].Hash)

	blockchian.AddBlockToBlockChain("send 100 RMB to Alia",blockchian.Blocks[len(blockchian.Blocks)-1].Height+1,blockchian.Blocks[len(blockchian.Blocks)-1].Hash)

	blockchian.AddBlockToBlockChain("send 100 RMB to Ada",blockchian.Blocks[len(blockchian.Blocks)-1].Height+1,blockchian.Blocks[len(blockchian.Blocks)-1].Hash)


	fmt.Println(blockchian.Blocks)

}
