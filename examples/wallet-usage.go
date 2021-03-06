package main

import (
	"fmt"
	"math/big"

	"github.com/jiuliangLiu/phoenixchain-go-sdk/common"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/network"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/web3go"
)

func main1() {
	const mnemonic = "always brick access science decade nasty marriage attack fame topple pen add"
	w, err := web3go.NewWalletByMnemonics(mnemonic)
	if err != nil {
		fmt.Println("import wall error: ", err)
		return
	}

	w.SetNetworkCfg(&network.DefaultMainNetConfig)

	accounts := w.Accounts()
	for _, account := range accounts {
		b, _ := w.BalanceOf(account.Address)
		addr, _ := account.ToMainNetAddress()
		fmt.Printf("balance of %s is %s\n", addr, b.String())
	}

	digest, err := w.Transfer(accounts[0].Address, common.MustStringToAddress("lat1u3vrx4n0hcmgdjzqk299xcvl2p5fqen8skh03c"), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("transfer failed: ", err)
		return
	}

	fmt.Println("tx send: ", digest)
}
