package main

import (
	"fmt"
	"bytechain/blockchian"
)

func main() {
	genesisBlock := blockchian.CreateGenesisBlock("创世区块")
	fmt.Println(genesisBlock)
}
