package main

import (
	"github.com/amidgo/amidtoken/routing"
	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-gonic/gin"
)

func main() {
	variables.Init()
	c := gin.Default()
	c.POST("/balance", routing.Balance)

	c.POST("/buyPrivate", routing.BuyPrivate)

	c.POST("/buyPublic", routing.BuyPublic)

	c.POST("/changeCost", routing.ChangeCost)

	c.GET("/time", routing.GetTime)

	c.GET("/phase", routing.GetPhase)

	c.GET("/cost", routing.GetCost)

	c.POST("/login", routing.Login)

	c.POST("/sendRequest", routing.SendRequest)

	c.GET("/requests", routing.Requests)

	c.POST("/handleRequest", routing.HandleRequest)

	c.POST("/timeTravel", routing.TimeTravel)

	c.POST("/transferPublic", routing.TransferPublic)

	c.POST("/transferPrivate", routing.TransferPrivate)

	c.GET("/users", routing.AllUsers)

	c.Run("0.0.0.0:1212")
}
