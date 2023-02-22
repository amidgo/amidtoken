package routing

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

type BuyBody struct {
	Sender
	Amount *big.Int `json:"amount"`
}

func Buy(ctx *gin.Context) {
	var body BuyBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	cost, _ := variables.Contract.Cost(variables.DefaultCallOpts())
	cost.Mul(cost, body.Amount)
	fmt.Println(cost)
	tOpts, _ := variables.TransactOpts(*body.Address, cost)
	_, err := variables.Contract.Buy(tOpts, body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
