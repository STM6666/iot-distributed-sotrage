package POW

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

var (
	//Nonce循环上限
	maxNonce = math.MaxInt64
)

// 难度值
const targetBits = 24

// pow结构
type proOfWork struct {
	block  *Block
	target *big.Int
}

func NewProfOfWork(b *Block) *proOfWork {
	//target为最终难度值
	target := big.NewInt(1)
	//target为1向左位移256-24(挖矿难度)
	target.Lsh(target, uint(256-targetBits))
	//生成pow结构
	pow := &proOfWork{b, target}
	return pow
}

// 挖矿与运行
func (pow *proOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("pow数据：%s,maxNonce%d\n", pow.block.Data, maxNonce)
	for nonce < maxNonce {
		//数据准备
		data := pow.prepareData(int64(nonce))
		//计算hash
		hash := sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		//按字节比较，hashInt.Cmp<0代表找到目标值Nonce
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]

}

// 准备函数，使用Join完成字节切片的组合
func (pow *proOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			Int2Hex(pow.block.Timestamp),
			Int2Hex(int64(targetBits)),
			Int2Hex(nonce),
		},
		[]byte{},
	)
	return data
}

// 组合时需要将Int转化为[]byte
func Int2Hex(num int64) []byte {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, num)
	return buff.Bytes()
}

// 校验区块正确性
func (pow *proOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
