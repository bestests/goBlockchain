package main

import (
	"fmt"
	"crypto/sha256"
	"time"
	"bytes"
	"strconv"
	"math/big"
	"math"
	"encoding/binary"
	"log"
	"os"
)

var (
	maxNonce = math.MaxInt64
	loger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
)

const targetBits = 24

type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

type Blockchain struct {
	blocks []*Block
}

type ProofOfWork struct {
	block *Block
	target *big.Int
}

// Block
func (b *Block) SetHash () {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock (data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock () *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Blockchain
func (bc *Blockchain) AddBlock (data string) {
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}


func NewBlockchain () *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// ProofOfWork
func NewProofOfWork (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData (nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run () (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	loger.Printf("Mining the block containing \"%s\"\n", pow.block.Data)


	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\rINFO: %s - %x : %d", nowStr(), hash, nonce);
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Println()
			break
		} else {
			nonce++
		}
	}


	/*
	data := pow.prepareData(nonce)
	hash = sha256.Sum256(data)
	fmt.Printf("\rINFO: %s - %x : %d", nowStr(), hash, nonce);
	hashInt.SetBytes(hash[:])
	fmt.Println()
	*/

	loger.Print("Mining End\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate () bool {

	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

// util
func IntToHex (num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func nowStr () (str string) {

	now := time.Now()
	str = now.Format("2006-01-02 15:04:05")

	return
}

func main () {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev Hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data      : %s\n", block.Data)
		fmt.Printf("Hash      : %x\n", block.Hash)
		fmt.Printf("Nonce     : %d\n", block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow       : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

