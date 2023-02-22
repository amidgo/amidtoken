package routing

import (
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

func TimeTravel(ctx *gin.Context) {
	_, err := variables.Contract.TimeTravel(variables.DefaultTransactOpts())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewRDataError(err))
		return
	}
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, NewRDataSuccess(nil))
}
