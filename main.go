package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}

func main() {
	one := []string{"User 1 sent 4 coins to BC"}
	abc := Blocks(one, []byte{})
	fmt.Println("This is the 1 block")
	fmt.Println(abc)

	two := []string{"User 2 sent 2 coins to ETH"}
	xyz := Blocks(two, []byte{})
	fmt.Println("This is the second block")
	fmt.Println(xyz)

}

func Blocks(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevHash,
		Hash:         NewHash(currentTime, transactions, prevHash),
	}
}

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	// This is a placeholder for a real hash function
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Show(block *Block) {
	fmt.Printf("\tTimestamp: %s\n", block.timestamp.String())

	fmt.Printf("\tPrevious Hash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %v\n", block.Hash)
	Transactions(block)

}

func Transactions(block *Block) {
	fmt.Println("\tTransactions: ")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}
