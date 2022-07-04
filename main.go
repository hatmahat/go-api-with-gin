package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// server side rendering (ssr) -> serving by backend
// api

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

	routerEngine := gin.Default()

	/*
		www.enigma.com/login
		www.enigma.com/register
		www.enigma.com/logout
		www.enigma.com

		grouping: master
		www.enigma.com/master/users
		www.enigma.com/master/student

		www.enigma.com/api/auth/login
	*/

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(200, "Healthy Check")
	})

	rgAuth := routerEngine.Group("/api/auth")
	rgAuth.POST("/login", login)

	rgMaster := routerEngine.Group("/api/master")
	rgMaster.GET("/greeting/:name", greeting)

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

/*
Model Binding & Validation
-> Shouldbind, ShouldbingJSON, -> using struct tag binding:required
*/

func login(c *gin.Context) {
	// username := c.PostForm("username")
	// c.PostForm("password")
	var uc UserCredential
	if err := c.BindJSON(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    uc.Username,
		})
	}
}
