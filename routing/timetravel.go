package routing

import (
	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

func TimeTravel(ctx *gin.Context) {
	variables.Contract.TimeTravel(variables.DefaultTransactOpts())
	RedirectFromRequestToRolePage(ctx)
}
