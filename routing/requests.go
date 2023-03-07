package routing

import (
	"math/big"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Sender *common.Address
	Name   string
}

func SendRequest(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Query("address"))
	name := ctx.Request.FormValue("name")
	tOpts, _ := variables.TransactOpts(address, big.NewInt(0))
	if _, err := variables.Contract.SendRequest(tOpts, name); err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}

func Requests() []*RequestBody {
	accs := make([]common.Address, 0)
	var index int64 = 0
	for {
		addr, err := variables.Contract.RequestAddresses(variables.DefaultCallOpts(), big.NewInt(index))
		index++
		if err != nil {
			break
		}
		accs = append(accs, addr)
	}

	requests := make([]*RequestBody, 0)
	for _, v := range accs {
		name, err := variables.Contract.Requests(variables.DefaultCallOpts(), v)
		if err != nil {
			continue
		}
		addr := v
		requests = append(requests, &RequestBody{Name: name, Sender: &addr})
	}
	return requests
}

func HandleRequest(ctx *gin.Context) {
	private := common.HexToAddress(ctx.Query("address"))
	address := common.HexToAddress(ctx.Request.FormValue("sender"))
	isAccept := ctx.Request.FormValue("status") == "on"
	tOpts, err := variables.TransactOpts(private, big.NewInt(0))
	_, err = variables.Contract.HandleRequest(tOpts, address, isAccept)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectFromRequestToRolePage(ctx)
}
