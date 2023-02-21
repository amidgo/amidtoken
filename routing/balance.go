package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type Sender struct {
	Address *common.Address `json:"address"`
}

type BalanceResponse struct {
	TokenBalance *big.Int `json:"tokenBalance"`
	EthBalance   *big.Int `json:"ethBalance"`
}

func Balance(ctx *gin.Context) {
	var sender Sender
	if err := ctx.BindJSON(&sender); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tokenBalance, err := variables.Contract.BalanceOf(variables.DefaultCallOpts(), *sender.Address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	blockNumber, err := variables.Client.BlockNumber(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ethBalance, err := variables.Client.BalanceAt(ctx, *sender.Address, big.NewInt(int64(blockNumber)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(&BalanceResponse{TokenBalance: tokenBalance, EthBalance: ethBalance}))
}
