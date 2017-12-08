package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

const CurlTest = `
curl http://localhost:7777/v1/AA_A
curl http://localhost:7777/v1/AAC_A
curl http://localhost:7777/v1/AA_B
curl http://localhost:7777/v1/AAB_A
curl http://localhost:7777/v1/AAB_B
curl http://localhost:7777/v1/AABC_A
curl http://localhost:7777/v1/AAB_C
`

func main() {
	r := gin.Default()
	APIv1 := r.Group("/v1")

	APIv1.Use(MiddlewareA)
	APIv1.Use(MiddlewareA) //这样User会连续调用两次Handler，因为本质是队列，所以不会覆盖
	APIv1.GET("AA_A", HandlerA) // MiddlewareA-> MiddlewareA-> HandlerA

	//这样用中间件而不是 APIv1.USE 可以确保不影响其他路由
	APIv1.GET("AAC_A", MiddlewareC, HandlerA)// MiddlewareA-> MiddlewareA-> MiddlewareC-> HandlerA
	APIv1.GET("AA_B",HandlerB)// MiddlewareA-> MiddlewareA-> HandlerB

	APIv1.Use(MiddlewareB)
	{
		//这里的括号有迷惑性，容易误认为括号外不会用到MiddlewareB
		//实际上APIv1在此时已装配两个MiddlewareA，和MiddlewareB，
		//所以之后即使在括号外 APIv1为基础的API一定会依次调用两个MiddlewareA，和MiddlewareB
		APIv1.GET("AAB_A", HandlerA)
	}
	//实际上MiddlewareB还在APIv1的handler 队列里
	APIv1.GET("AAB_B", HandlerB)

	//真正意义上的作用域可能要这么写，AAC_A的方法对单个路由更方便
	{
		//这里利用Group new了一个RouteGroup，所以不影响APIv1
		//这里tmpAPI对性能不会带来很大影响，因为APIv1，tmpAPI都只是注册Handler，用完就不需要了
		tmpAPI := APIv1.Group("").Use(MiddlewareC)
		tmpAPI.GET("AABC_A", HandlerA)
	}
	//对APIv1没有影响，外部也调用不掉tmpAPI
	APIv1.GET("AAB_C", HandlerC)

	r.Run(":7777")
}

func MiddlewareA(c *gin.Context) {
	fmt.Println("MiddlewareA")
	c.Next()
}
func MiddlewareB(c *gin.Context) {
	fmt.Println("MiddlewareB")
	c.Next()
}
func MiddlewareC(c *gin.Context) {
	fmt.Println("MiddlewareC")
	c.Next()
}
func HandlerA(c *gin.Context) {
	fmt.Println("HandlerA")
	c.JSON(200, gin.H{"status": "OK"})
}
func HandlerB(c *gin.Context) {
	fmt.Println("HandlerB")
	c.JSON(200, gin.H{"status": "OK"})
}
func HandlerC(c *gin.Context) {
	fmt.Println("HandlerC")
	c.JSON(200, gin.H{"status": "OK"})
}
