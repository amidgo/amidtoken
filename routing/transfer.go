package routing

import (
	"math/big"
	"strconv"

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
	address := common.HexToAddress(ctx.Query("address"))
	to := common.HexToAddress(ctx.Request.FormValue("to"))
	rAmount, err := strconv.ParseInt(ctx.Request.FormValue("amount"), 10, 64)
	amount := big.NewInt(rAmount)
	tOpts, _ := variables.TransactOpts(address, big.NewInt(0))
	_, err = variables.Contract.Transfer(tOpts, to, amount)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}

func TransferFrom(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Query("address"))
	from := common.HexToAddress(ctx.Request.FormValue("from"))
	to := common.HexToAddress(ctx.Request.FormValue("to"))
	rAmount, err := strconv.ParseInt(ctx.Request.FormValue("amount"), 10, 64)
	amount := big.NewInt(rAmount)
	tOpts, err := variables.TransactOpts(address, big.NewInt(0))
	_, err = variables.Contract.TransferFrom(tOpts, from, to, amount)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}
