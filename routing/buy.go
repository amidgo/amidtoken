package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

type BuyBody struct {
	Sender
	Amount *big.Int `json:"amount"`
	Value  *big.Int `json:"value"`
}

func Buy(ctx *gin.Context) {
	var body BuyBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tOpts, _ := variables.TransactOpts(*body.Address, big.NewInt(0))
	_, err := variables.Contract.Buy(tOpts, body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
