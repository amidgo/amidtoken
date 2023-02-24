package routing

import (
	"math/big"
	"net/http"
	"time"

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
	SystemTime *big.Int `json:"systemTime"`
	TimeNow    int64    `json:"timeNow"`
	TimeStart  *big.Int `json:"timeStart"`
}

func GetTime(ctx *gin.Context) {
	sTime, err := variables.Contract.GetTime(variables.DefaultCallOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	timeStart, err := variables.Contract.StartTime(variables.DefaultCallOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	timeNow := time.Now().Unix()
	ctx.JSON(http.StatusOK, NewRDataSuccess(&TimeBody{SystemTime: sTime, TimeStart: timeStart, TimeNow: timeNow}))
}
