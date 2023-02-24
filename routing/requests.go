package routing

import (
	"math/big"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type SendRequestBody struct {
	Sender
	Name string `json:"name"`
}

func SendRequest(ctx *gin.Context) {
	var body SendRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	tOpts, _ := variables.TransactOpts(*body.Address, big.NewInt(0))
	if _, err := variables.Contract.SendRequest(tOpts, body.Name); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}

func Requests(ctx *gin.Context) {
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

	requests := make([]*SendRequestBody, 0)
	for _, v := range accs {
		name, err := variables.Contract.Requests(variables.DefaultCallOpts(), v)
		if err != nil {
			continue
		}
		addr := v
		requests = append(requests, &SendRequestBody{Name: name, Sender: Sender{&addr}})
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(requests))
}

type HandleRequestBody struct {
	Sender
	IsAccept bool `json:"isAccept"`
}

func HandleRequest(ctx *gin.Context) {
	var body HandleRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	_, err := variables.Contract.HandleRequest(variables.DefaultTransactOpts(), *body.Address, body.IsAccept)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
