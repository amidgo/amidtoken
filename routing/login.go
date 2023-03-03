package routing

import (
	"errors"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Sender
	Password string `json:"password"`
}

type LoginResponse struct {
	Role string `json:"role"`
}

type RData struct {
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Success bool        `json:"success"`
}

func NewRDataError(err error) *RData {
	return &RData{nil, err.Error(), false}
}

func NewRDataSuccess(data interface{}) *RData {
	return &RData{data, "nil", true}
}

func RedirectToError(ctx *gin.Context, err error) {
	ctx.Redirect(http.StatusMovedPermanently, "/error?err="+err.Error())
}

func RedirectToRolePage(ctx *gin.Context, addr string, role string) {
	ctx.Redirect(http.StatusMovedPermanently, "/user-page"+"?address="+addr+"&role="+role)
}

func RedirectFromRequestToRolePage(ctx *gin.Context) {
	role := ctx.Query("role")
	address := ctx.Query("address")
	time.Sleep(time.Microsecond * 500)
	RedirectToRolePage(ctx, address, role)
}

func Login(ctx *gin.Context) {
	addr := common.HexToAddress(ctx.Request.FormValue("login"))
	password := ctx.Request.FormValue("password")
	role, err := variables.Contract.Users(variables.DefaultCallOpts(), addr)
	if err != nil {
		RedirectToError(ctx, err)
		return
	}
	acc := variables.ImportAccount(addr)
	if acc == nil {
		RedirectToError(ctx, errors.New("wrong login"))
		return
	}
	if err := variables.Keystore.Unlock(*acc, password); err != nil {
		RedirectToError(ctx, err)
		return
	}
	RedirectToRolePage(ctx, addr.String(), role)
}
