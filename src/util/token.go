package util

import (
	"time"
	"fmt"
	"crypto/md5"
	"strconv"
	"io"
)

func SetToken(username string) (string) {
	crutime := time.Now().Unix()
	//fmt.Println("crutime-->", crutime)
	
	h := md5.New()
	//fmt.Println("h-->", h)
	
	//fmt.Println("strconv.FormatInt(crutime, 10)-->", strconv.FormatInt(crutime, 10))
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	
	//fmt.Println("h-->", h)
	
	token := fmt.Sprintf("%x", h.Sum(nil))
	//fmt.Println("token--->", token)
	//fmt.Println(len(token))
	
	return username+"_"+token
}
