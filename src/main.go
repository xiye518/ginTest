package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	
	"API"
)

func main() {
	
	r := gin.Default()
	
	v1 := r.Group("api/v1")
	{
		v1.POST("/user", API.PostUser)
		v1.GET("/user", API.GetUsers)
		v1.GET("/user/:id", API.GetUser)
		v1.PUT("/user/:id", API.UpdateUser)
		v1.DELETE("/user/:id", API.DeleteUser)
	}
	
	r.Run(":8080")
}
