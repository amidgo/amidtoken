package routing

import (
	"math/big"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TransferBody struct {
	Sender
	To     *common.Address `json:"to"`
	Amount *big.Int        `json:"amount"`
}

type TransferFromBody struct {
	From   *common.Address `json:"from"`
	To     *common.Address `json:"to"`
	Amount *big.Int        `json:"amount"`
}

func Transfer(ctx *gin.Context) {
	var body TransferBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tOpts, _ := variables.TransactOpts(*body.Address, big.NewInt(0))
	_, err := variables.Contract.Transfer(tOpts, *body.To, body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}

func TransferFrom(ctx *gin.Context) {
	var body TransferFromBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	_, err := variables.Contract.TransferFrom(variables.DefaultTransactOpts(), *body.From, *body.To, body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
