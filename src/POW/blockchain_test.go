package POW

import (
	"fmt"
	"testing"
)

func TestBlockchain(t *testing.T) {
	//初始化
	bc := NewBlockchain()
	//创建块记录
	bc.AddBlock("helloworld")

	//遍历
	for _, block := range bc.Blocks {
		fmt.Printf("prev,hash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("nonce:%d\n", block.Nonce)
		pow := NewProfOfWork(block)
		fmt.Printf("Pow: %t\n", pow.Validate())
		fmt.Println()
		fmt.Printf("难度值:%d\n", int64(targetBits))

	}
}
