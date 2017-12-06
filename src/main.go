package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	
	"api"
)

func main() {
	api.Debug = true
	r := gin.Default()
	
	//authrization:=r.Group("/",jwtauth中间件)
	//{
	//	authrization.Get("user",handleFunc)
	//}
	
	v1 := r.Group("api/v1")
	{
		v1.POST("/login", api.LoginApi)
		v1.POST("/reg", api.RegisterApi)
		v1.GET("/reg", api.RegisterApi)
		v1.GET("/show", api.ShowAllApi)
		
		//v1.POST("/user", API.PostUser)
		//v1.GET("/user", API.GetUsers)
		//v1.GET("/user/:id", API.GetUser)
		//v1.PUT("/user/:id", API.UpdateUser)
		//v1.DELETE("/user/:id", API.DeleteUser)
	}
	
	r.Run(":8080")
}
