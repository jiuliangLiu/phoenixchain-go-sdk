package network

import "math/big"

const (
	MainNetHrp = ""
	TestNetHrp = ""
)

var (
	// TODO(liangqin.fan): 主网上线后, 需要更新为主网的真实chainid
	MainNetChainId = big.NewInt(210309)
	TestNetChainId = big.NewInt(1025)
)

type Config struct {
	Url     string
	ChainId *big.Int
}

var (
	// TODO(liangqin.fan): 主网上线后, 需要更新为主网的正式RPC地址
	DefaultTestNetConfig = Config{"http://39.104.68.32:6789", TestNetChainId}
	DefaultMainNetConfig = Config{"http://35.247.155.162:6789", MainNetChainId}
	//DefaultMainNetConfig = Config{"http://172.16.64.132:6789", MainNetChainId}
)

type PposNetworkParameters struct {
	Config
	Hrp                     string
	RestrictingPlanContract string
	StakingContract         string
	SlashContract           string
	ProposalContract        string
	RewardContract          string
}

var (
	PposMainNetParams = &PposNetworkParameters{
		DefaultMainNetConfig,
		MainNetHrp,
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzsjx8h7",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyva9ztf",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq93t3hkm",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxlcypcy",
	}

	PposTestNetParams = &PposNetworkParameters{
		DefaultTestNetConfig,
		TestNetHrp,
		"0x1000000000000000000000000000000000000001",
		"0x1000000000000000000000000000000000000002",
		"0x1000000000000000000000000000000000000004",
		"0x1000000000000000000000000000000000000005",
		"0x1000000000000000000000000000000000000003",
	}
)
