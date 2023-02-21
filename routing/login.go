package routing

import (
	"net/http"

	"github.com/amidgo/amidtoken/variables"
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

func Login(ctx *gin.Context) {
	var body LoginBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	role, err := variables.Contract.Users(variables.DefaultCallOpts(), *body.Address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	acc := variables.ImportAccount(*body.Address)
	if err := variables.Keystore.Unlock(*acc, body.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(&LoginResponse{role}))
}
