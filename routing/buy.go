package routing

import (
	"math/big"
	"strconv"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type BuyBody struct {
	Sender
	Amount *big.Int `json:"amount"`
}

func Buy(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Query("address"))
	number, err := strconv.ParseInt(ctx.Request.FormValue("amount"), 10, 64)
	amount := big.NewInt(number)
	cost, _ := variables.Contract.Cost(variables.DefaultCallOpts())
	cost.Mul(cost, amount)
	tOpts, _ := variables.TransactOpts(address, cost)
	_, err = variables.Contract.Buy(tOpts, amount)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}
