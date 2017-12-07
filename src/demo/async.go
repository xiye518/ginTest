package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"time"
	"log"
	"net/http"
)

func main(){
	version1()
	//version2()
	//version3()
}

func version1(){
	router:=gin.Default()
	
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
		c.JSON(200, gin.H{"msg": "Done! in path" + c.Request.URL.Path})
	})
	//curl -i Get http://localhost:8000/sync
	
	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
			c.JSON(200, gin.H{"msg": "Done! in path" + cCp.Request.URL.Path})
		}()
	})
	
	router.Run(":8000")
}

func version2(){
	router:=gin.Default()
	
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
		c.JSON(200, gin.H{"msg": "Done! in path" + c.Request.URL.Path})
	})
	//curl -i Get http://localhost:8000/sync
	
	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
			c.JSON(200, gin.H{"msg": "Done! in path" + cCp.Request.URL.Path})
		}()
	})
	
	http.ListenAndServe(":8000",router)
}

func version3(){
	router:=gin.Default()
	
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
		c.JSON(200, gin.H{"msg": "Done! in path" + c.Request.URL.Path})
	})
	//curl -i Get http://localhost:8000/sync
	
	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
			c.JSON(200, gin.H{"msg": "Done! in path" + cCp.Request.URL.Path})
		}()
	})
	
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}