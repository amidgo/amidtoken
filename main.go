package main

import (
	"github.com/amidgo/amidtoken/routing"
	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	variables.Init()
	c := gin.Default()
	c.Use(cors.Default())

	c.POST("/approve", routing.Approve)

	c.POST("/balance", routing.Balance)

	c.POST("/buy", routing.Buy)

	c.POST("/changeCost", routing.ChangeCost)

	c.GET("/time", routing.GetTime)

	c.GET("/phase", routing.GetPhase)

	c.GET("/cost", routing.GetCost)

	c.POST("/login", routing.Login)

	c.POST("/sendRequest", routing.SendRequest)

	c.GET("/requests", routing.Requests)

	c.POST("/handleRequest", routing.HandleRequest)

	c.POST("/timeTravel", routing.TimeTravel)

	c.POST("/transfer", routing.Transfer)

	c.POST("/transferFrom", routing.TransferFrom)

	c.GET("/users", routing.AllUsers)

	c.Run("0.0.0.0:1212")
}
