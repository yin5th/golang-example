package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

//定义一个区块的结构
//一个区块应当包括最基本的 区块数据、区块时间、区块的hash、上一个区块的hash【便于有序串联】
type Block struct {
	Timestamp     int64  //时间戳
	Data          []byte //当前区块的信息（如账单、操作记录等等）
	PrevBlockHash []byte //上一个区块的哈希
	Hash          []byte //当前hash
}

//生成当前区块的hash
func (this *Block) SetHash() {
	//将本区块的Timestamp|Data|PrevBlockHash进行拼接并且hash处理
	//1. Timestamp转化成[]byte
	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10))
	//2. 拼接 （以空二进制进行拼接）
	headers := bytes.Join([][]byte{this.PrevBlockHash, this.Data, timestamp}, []byte{})
	//3. 将headers进行sha256加密
	hash := sha256.Sum256(headers)

	this.Hash = hash[:]
}

//创建一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	//生成一个空的区块
	block := Block{}
	//获取当前时间戳
	block.Timestamp = time.Now().Unix()

	//上一个区块的hash
	block.PrevBlockHash = prevBlockHash

	//获取当前区块的数据
	block.Data = []byte(data)

	//当前区块的hash 先设为空
	block.Hash = []byte{}

	//生成当前区块的hash
	block.SetHash()

	return &block
}

//定义区块链的结构
type BlockChain struct {
	Blocks []*Block
}

//向区块链添加一个区块
func (this *BlockChain) AddBlock(data string) {
	//找到新区块的前一个区块
	prevBlock := this.Blocks[len(this.Blocks)-1]
	//生成新区块
	newBlock := NewBlock(data, prevBlock.Hash)

	//追加新区块到区块链中
	this.Blocks = append(this.Blocks, newBlock)
}

//区块链= 创世块-->区块-->区块
//新建创世块 【即每个区块链中的第一个区块】
func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("Genesis block")
	genesisBlock.PrevBlockHash = []byte{}
	return &genesisBlock
}

//新建一个区块链
func NewBlockChain() *BlockChain {
	//新建一个区块链 并给予创世块
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	//新建区块链
	bc := NewBlockChain()

	for {
		//用户输入的指令 1,2,其他
		var cmd string
		fmt.Println("1 添加区块")
		fmt.Println("2 查看所有区块")
		fmt.Println("其他任意键退出")
		//接收键盘输入
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case "1":
			input := make([]byte, 1024)
			fmt.Println("请输入区块数据：")
			os.Stdin.Read(input)
			bc.AddBlock(string(input))
			fmt.Println("区块添加成功！")
		case "2":
			for i, block := range bc.Blocks {
				fmt.Println("========================")
				fmt.Println("第 ", i, " 个 区块的信息：")
				fmt.Printf("PrevHash： %x\n", block.PrevBlockHash)
				fmt.Printf("Data： %s\n", block.Data)
				fmt.Printf("Hash： %x\n", block.Hash)
				fmt.Println("========================")
			}
		default:
			//退出程序
			fmt.Println("bye!")
			return
		}
	}
}
