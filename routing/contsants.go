package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

type CostBody struct {
	Cost *big.Int `json:"cost"`
}

func GetCost(ctx *gin.Context) {
	cost, err := variables.Contract.Cost(variables.DefaultCallOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(&CostBody{Cost: cost}))
}

type PhaseBody struct {
	Phase string `json:"phase"`
}

func GetPhase(ctx *gin.Context) {
	phase, err := variables.Contract.GetPhase(variables.DefaultCallOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(&PhaseBody{Phase: phase}))
}

type TimeBody struct {
	Time *big.Int `json:"time"`
}

func GetTime(ctx *gin.Context) {
	time, err := variables.Contract.GetTime(variables.DefaultCallOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(&TimeBody{Time: time}))
}
