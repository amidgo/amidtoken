package routing

import (
	"math/big"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

type ChangeCostBody struct {
	NewValue *big.Int `json:"newValue"`
}

func ChangeCost(ctx *gin.Context) {
	var body ChangeCostBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tOpts := variables.DefaultTransactOpts()
	_, err := variables.Contract.ChangeCost(tOpts, body.NewValue)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
