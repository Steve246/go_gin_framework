package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	routerEngine := gin.Default()
	routerEngine.GET("/ping-pong", func(c *gin.Context) {
		c.String(200, "Healthy Check")
	})
	err := routerEngine.Run(":8888") //default, kalau kosong pake port 8080
	if err != nil {
		panic(err)
	}
}
