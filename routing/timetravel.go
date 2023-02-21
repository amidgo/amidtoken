package routing

import (
	"net/http"

	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

func TimeTravel(ctx *gin.Context) {
	_, err := variables.Contract.TimeTravel(variables.DefaultTransactOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
