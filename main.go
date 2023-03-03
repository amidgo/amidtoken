package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amidgo/amidtoken/routing"
	"github.com/amidgo/amidtoken/variables"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	time.Sleep(time.Second)
	variables.Init()
	c := gin.Default()
	c.Use(cors.Default())

	c.LoadHTMLGlob("templates/*")

	c.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	c.POST("/login", routing.Login)

	c.GET("/error", func(ctx *gin.Context) {
		err := ctx.Query("err")
		ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": err})
	})

	c.GET("/user-page", routing.UserPage)

	c.POST("/approve", routing.Approve)
	c.POST("/buy", routing.Buy)
	c.POST("/sendRequest", routing.SendRequest)
	c.POST("/transfer", routing.Transfer)
	c.POST("/transferFrom", routing.TransferFrom)
	c.POST("/timeTravel", routing.TimeTravel)
	c.POST("/changeCost", routing.ChangeCost)
	c.POST("/handleRequest", routing.HandleRequest)

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
