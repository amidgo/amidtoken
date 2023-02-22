package routing

import (
	"math/big"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ApproveBody struct {
	Sender
	To    *common.Address `json:"to"`
	Value *big.Int        `json:"value"`
}

func Approve(ctx *gin.Context) {
	var body ApproveBody
	ctx.BindJSON(&body)
	tOpts, err := variables.TransactOpts(*body.Address, big.NewInt(0))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	_, err = variables.Contract.Approve(tOpts, tOpts.From, *body.To, body.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
