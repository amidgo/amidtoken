package variables

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func NewCallOpts(from common.Address) *bind.CallOpts {
	blockNumber, _ := Client.BlockNumber(context.Background())
	return &bind.CallOpts{Pending: true, From: from, BlockNumber: big.NewInt(int64(blockNumber)), Context: context.Background()}
}

func DefaultCallOpts() *bind.CallOpts {
	return NewCallOpts(Owner)
}
