package typedefs

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/jiuliangLiu/phoenixchain-go-sdk/common"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/core/types"
	"github.com/jiuliangLiu/phoenixchain-go-sdk/crypto"
)

type Credentials struct {
	privateKey *ecdsa.PrivateKey
	hrp        string
}

func NewCredential(hexPrivateKeyStr string, hrp string) (*Credentials, error) {
	if len(hexPrivateKeyStr) < 64 {
		return nil, fmt.Errorf("length of private key is not fullfilled")
	}

	if strings.HasPrefix(hexPrivateKeyStr, "0x") || strings.HasPrefix(hexPrivateKeyStr, "0X") {
		hexPrivateKeyStr = hexPrivateKeyStr[2:]
	}

	key, err := crypto.HexToECDSA(hexPrivateKeyStr)
	if err != nil {
		return nil, err
	}

	return &Credentials{key, hrp}, nil
}

func (c *Credentials) Address() common.Address {
	return crypto.PubkeyToAddress(c.privateKey.PublicKey)
}

func (c *Credentials) HexAddress() string {
	addr := c.Address()
	return addr.String()
}

func (c *Credentials) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, c.privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (c *Credentials) MustStringToAddress(addr string) common.Address {
	return common.MustStringToAddress(addr)
}

func (c *Credentials) PrivateKey() *ecdsa.PrivateKey {
	return c.privateKey
}
