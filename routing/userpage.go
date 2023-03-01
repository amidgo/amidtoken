package routing

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type UserStruct struct {
	Address *common.Address
	Role    string
	Eth     *big.Int
	CMON    *big.Int
}

func UserPage(ctx *gin.Context) {

}
