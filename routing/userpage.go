package routing

import (
	"math/big"
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type UserStruct struct {
	Address *common.Address
	Role    string
	Eth     *big.Int
	CMON    *big.Int
}

type TimeStruct struct {
	Start   uint
	Private uint
	Public  uint
}

type AdminUserInfo struct {
	Address        *common.Address
	EthBalance     *big.Int
	PrivateBalance *big.Int
	PublicBalance  *big.Int
	TotalBalance   *big.Int
}

func UserPage(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Query("address"))
	role := ctx.Query("role")
	blockNumber, _ := variables.Client.BlockNumber(ctx)
	eth, _ := variables.Client.BalanceAt(ctx, address, big.NewInt(int64(blockNumber)))
	cmon, _ := variables.Contract.BalanceOf(variables.DefaultCallOpts(), address)
	t, _ := variables.Contract.StartTime(variables.DefaultCallOpts())
	ct, _ := variables.Contract.GetTime(variables.DefaultCallOpts())
	currentTime := ct.Int64()
	startTime := t.Int64()
	privateTime := currentTime - 5
	publicTime := currentTime - 15
	users := AllUsers()
	requests := Requests()
	if privateTime < 0 {
		privateTime = 0
	}
	if publicTime < 0 {
		publicTime = 0
	}
	templateDict := map[string]string{
		"user":    "user.html",
		"public":  "public.html",
		"private": "private.html",
		"owner":   "owner.html",
	}

	ctx.HTML(http.StatusOK, templateDict[role], gin.H{
		"UserData": &UserStruct{Address: &address, Role: role, Eth: eth, CMON: cmon},
		"TimeData": &TimeStruct{Start: uint(startTime), Private: uint(privateTime), Public: uint(publicTime)},
		"Users":    users,
		"Requests": requests,
	})

}
