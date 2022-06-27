package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/jiuliangLiu/phoenixchain-go-sdk/network"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/ppos"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/ppos/typedefs"
)

func TestGetBlockReward(t *testing.T) {
	const PrivateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"
	var credentials, _ = typedefs.NewCredential(PrivateKey, network.TestNetHrp)

	config := network.PposTestNetParams
	sc := ppos.NewStakingContract(config, credentials)

	list, err := sc.GetPackageReward()
	if err != nil {
		fmt.Println("StakingContract.GetPackageReward failed: ", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		fmt.Println("Marshal of list failed: ", err)
	}

	fmt.Println(string(result))
}

func TestGetStakingReward(t *testing.T) {
	const PrivateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"
	var credentials, _ = typedefs.NewCredential(PrivateKey, network.TestNetHrp)

	config := network.PposTestNetParams
	sc := ppos.NewStakingContract(config, credentials)

	list, err := sc.GetStakingReward()

	if err != nil {
		t.Errorf("StakingContract.GetStakingReward failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}
	fmt.Println(string(result))
}

func TestGetDelegateReward(t *testing.T) {
	const PrivateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"
	var credentials, _ = typedefs.NewCredential(PrivateKey, network.TestNetHrp)

	config := network.PposMainNetParams
	dc := ppos.NewDelegateContract(config, credentials)

	nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"

	amount := new(big.Int)
	amount.SetString("200000000000000000000", 10)
	list, err := dc.Delegate(nodeId, typedefs.FREE_AMOUNT_TYPE, amount)
	if err != nil {
		t.Errorf("DelegateContract.Delegate failed: %s", err)
	}

	result, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Marshal of list failed: %s", err)
	}

	fmt.Println(string(result))
}
