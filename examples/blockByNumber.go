package main

import (
	"fmt"

	"github.com/jiuliangLiu/phoenixchain-go-sdk/web3go"
)

func main() {
	const alayaEndpoint = "http://39.104.62.41:6790"
	web3g, _ := web3go.New(alayaEndpoint)
	// a := big.NewInt(890)
	var test int64
	test = 226955
	block, err := web3g.BlockByNumber(test)
	fmt.Println(block, err)
}
