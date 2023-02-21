package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type Sender struct {
	Address *common.Address `json:"address"`
}

type BalanceResponse struct {
	PublicBalance  *big.Int `json:"publicBalance"`
	PrivateBalance *big.Int `json:"privateBalance"`
}

func Balance(ctx *gin.Context) {
	var sender Sender
	if err := ctx.BindJSON(&sender); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	st, err := variables.Contract.BalanceOf(variables.DefaultCallOpts(), *sender.Address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(BalanceResponse{st.PublicTokens, st.PrivateTokens}))
}
