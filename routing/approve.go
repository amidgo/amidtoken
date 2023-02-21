package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ApproveBody struct {
	Sender
	To    *common.Address `json:"to"`
	Value *big.Int        `json:"value"`
}

func ApprovePublic(ctx *gin.Context) {
	var body ApproveBody
	ctx.BindJSON(&body)
	tOpts, err := variables.TransactOpts(*body.Address, big.NewInt(0))
	_, err = variables.Contract.ApprovePublic(tOpts, tOpts.From, *body.To, body.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}

func ApprovePrivate(ctx *gin.Context) {
	var body ApproveBody
	ctx.BindJSON(&body)
	tOpts, err := variables.TransactOpts(*body.Address, big.NewInt(0))
	_, err = variables.Contract.ApprovePrivate(tOpts, tOpts.From, *body.To, body.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
