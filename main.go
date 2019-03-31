package main

import (
	"bytechain/blockchian"
)

func main() {
	blockchian := blockchian.CreateBlockchainWithGenesisBlock()

	defer blockchian.DB.Close()

	blockchian.AddBlockToBlockchain("send 100 RMB to TOM")

	blockchian.AddBlockToBlockchain("send 100 RMB to Bob")

	blockchian.AddBlockToBlockchain("send 100 RMB to Alia")

	blockchian.AddBlockToBlockchain("send 100 RMB to Ada")

	blockchian.Printchain()

}