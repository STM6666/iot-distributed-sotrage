package POW

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //时间戳
	Data          []byte //数据域
	PrevBlockHash []byte //前一区块hash值
	Hash          []byte //当前区块hash
	Nonce         int64
}

// 区块设置内部hash方法
func (b *Block) SetHash() {
	//将时间戳转换为[]byte
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	//将前一区块的hash、交易信息、时间戳联合到一起
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	//计算本块hash值
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 创建Block，返回Block指针
func NewBlock(data string, PrevBlockHash []byte) *Block {
	//先构造block
	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}, 0}
	//需要先挖矿

	pow := NewProfOfWork(block)
	nonce, hash := pow.Run()
	//设置hash和nonce
	block.Hash = hash
	block.Nonce = int64(nonce)
	return block
}

// 创世块创建，返回创世块Block指针
func NewGenesisBlock() *Block {
	return NewBlock("创世块", []byte{})

}
