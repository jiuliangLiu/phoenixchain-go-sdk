package resp

import (
	"github.com/jiuliangLiu/phoenixchain-go-sdk/common/hexutil"
)

type Reward struct {
	NodeId     string       `json:"nodeId"`
	StakingNum int          `json:"stakingNum"`
	Reward     *hexutil.Big `json:"reward"`
}
