package routing

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func AllUsers(ctx *gin.Context) {
	accs := make([]common.Address, 0)
	var err error
	var addr common.Address
	var index int64
	fmt.Println(index)
	for err == nil {
		addr, err = variables.Contract.UserAddresses(variables.DefaultCallOpts(), big.NewInt(index))
		index++
		if err != nil {
			continue
		}
		fmt.Println(addr)
		accs = append(accs, addr)
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(accs))
}
