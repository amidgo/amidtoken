package routing

import (
	"math/big"
	"strconv"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ChangeCostBody struct {
	NewValue *big.Int `json:"newValue"`
}

func ChangeCost(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Query("address"))
	v := ctx.Request.FormValue("value")
	value, _ := strconv.ParseInt(v, 10, 64)
	newCost := big.NewInt(value)
	tOpts, _ := variables.TransactOpts(address, big.NewInt(0))
	_, err := variables.Contract.ChangeCost(tOpts, newCost)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}
