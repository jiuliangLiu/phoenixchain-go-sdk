package main

import (
	"fmt"
	"testing"

	"github.com/jiuliangLiu/phoenixchain-go-sdk/common"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/web3go"
)

func TestGetTranR(t *testing.T) {
	web3g, _ := web3go.New("http://39.104.68.32:6790")
	tx, err := web3g.TransactionReceipt(common.HexToHash("0xceb52507488d08a90fb636d69e5d308d1dc55306f9ae2877c0ae1598249ea69c"))
	// tx.MarshalJSON()
	result, err := tx.MarshalJSON()

	fmt.Println(string(result), err)
}
