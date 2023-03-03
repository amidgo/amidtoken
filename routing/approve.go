package routing

import (
	"math/big"
	"strconv"

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
	address := common.HexToAddress(ctx.Query("address"))
	to := common.HexToAddress(ctx.Request.FormValue("to"))
	v, err := strconv.ParseInt(ctx.Request.FormValue("value"), 10, 64)
	value := big.NewInt(v)
	tOpts, err := variables.TransactOpts(address, big.NewInt(0))
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	_, err = variables.Contract.Approve(tOpts, to, value)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}
