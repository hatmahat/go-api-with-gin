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

	routerEngine.GET("/greeting/:name", greeting)

	err := routerEngine.Run("localhost:8888")
	if err != nil {
		panic(err)
	}

}

func greeting(c *gin.Context) {
	name := c.Param("name") // with parameter :name at path /greeting
	kec := c.Query("kec")   //Query param / String in path -> key?=value
	kel := c.Query("kel")

	/*
		Any required or mandatory attributes should be added as path param
		Any optional attributes should be added as query param
		params used for filtering data are usually used as query param
	*/

	c.String(200, "Greeting path... %s %s %s", name, kec, kel)
}
