package main

import (
	"fmt"
	"bytechain/blockchian"
)

func main() {
	block := blockchian.NewBlock("Genenis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(block)
}
