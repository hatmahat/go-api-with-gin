package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

	routerEngine := gin.Default()
	routerEngine.GET("/", func(c *gin.Context) {
		c.String(200, "Healthy Check")
	})
	err := routerEngine.Run(":8888")
	if err != nil {
		panic(err)
	}

}
