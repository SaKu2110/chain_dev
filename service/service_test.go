package service

import(
	"fmt"
	"time"
	"testing"
	"github.com/SaKu2110/oauth_chain/model"
)

func TestSuccess(t *testing.T) {
	data := model.Data{
		Amount: 1000,
	}
	block := Block{
		Index: 0,
		previousHash: "Genesis block",
		Timestamp: time.Now().Add(time.Hour * 72).Unix(),
		Data: data,
	}
	hash, _ := block.CreateHash()
	fmt.Printf("%x\n", hash)
	coin := BlockChain{
		Hash: hash,
		block: block,
	}
	// ---
	data = model.Data{
		Amount: 100,
	}
	block = Block{
		Index: 1,
		Timestamp: time.Now().Add(time.Hour * 72).Unix(),
		Data: data,
	}
	_ =coin.AddBlock(block)
	fmt.Printf("%x\n", coin.Hash)
	// ---
	data = model.Data{
		Amount: 10,
	}
	block = Block{
		Index: 2,
		Timestamp: time.Now().Add(time.Hour * 72).Unix(),
		Data: data,
	}
	_ =coin.AddBlock(block)
	fmt.Printf("%x\n", coin.Hash)
}