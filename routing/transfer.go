package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TransferBody struct {
	Sender
	To     *common.Address `json:"to"`
	Amount *big.Int        `json:"amount"`
}

func TransferPublic(ctx *gin.Context) {
	var body TransferBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tOpts, _ := variables.TransactOpts(*body.Address, big.NewInt(0))
	_, err := variables.Contract.TransferPublic(tOpts, *body.To, body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}

func TransferPrivate(ctx *gin.Context) {
	var body TransferBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tOpts, _ := variables.TransactOpts(*body.Address, big.NewInt(0))
	_, err := variables.Contract.TransferPrivate(tOpts, *body.To, body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
