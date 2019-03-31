package main

import (
	"bytechain/blockchian"
)

func main() {
	blockchian := blockchian.CreateBlockChainWithGenesisBlock()

	//new block

	blockchian.AddBlockToBlockChain("send 100 RMB to TOM")

	blockchian.AddBlockToBlockChain("send 100 RMB to Bob")

	blockchian.AddBlockToBlockChain("send 100 RMB to Alia")

	blockchian.AddBlockToBlockChain("send 100 RMB to Ada")
}