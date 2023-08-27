package POW

// 区块链：一个区块的指针切片
type Blockcahin struct {
	Blocks []*Block
}

// 为BlockChain增加AddBlock方法（向切片增加一个指针）
// 增加区块
func (bc Blockcahin) AddBlock(data string) {
	//获取前一块信息
	PrevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, PrevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 准备NewBlockchain函数，通过创世块创建并初始化区块链。
func NewBlockchain() *Blockcahin {
	return &Blockcahin{[]*Block{NewGenesisBlock()}}
}
