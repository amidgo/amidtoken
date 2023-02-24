package main

import (
	"fmt"
	"time"

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

	go func() {
		for {
			time.Sleep(time.Second)
			sTime, _ := variables.Contract.GetTime(variables.DefaultCallOpts())
			timeI := sTime.Int64()
			if timeI >= 5 && timeI < 15 {
				if _, err := variables.Contract.SetPrivatePhase(variables.DefaultTransactOpts()); err != nil {
					fmt.Println("PRIVATE")
					fmt.Println(err)
				}
				time.Sleep(time.Second)
			}
			if timeI >= 15 {
				if _, err := variables.Contract.SetPublicPhase(variables.DefaultTransactOpts()); err != nil {
					fmt.Println("PUBLIC")
					fmt.Println(err)
				}
				time.Sleep(time.Second)
			}
		}
	}()

	c.Run("0.0.0.0:1212")

}
