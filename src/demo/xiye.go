package main

import (
	"fmt"
	"time"
	"crypto/md5"
	"strconv"
	"io"
)

func main(){
	crutime := time.Now().Unix()
	fmt.Println( crutime)
	fmt.Println(int(crutime))
	fmt.Println(fmt.Sprintf("%T", crutime))
	
	h := md5.New()
	//fmt.Println("h-->", h)
	
	//fmt.Println("strconv.FormatInt(crutime, 10)-->", strconv.FormatInt(crutime, 10))
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	
	//fmt.Println("h-->", h)
	
	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("token--->", token)
	fmt.Println(len(token))
}

