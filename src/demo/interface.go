package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	fmt.Println(count(10))
	
	//testInterface()
}

func count(i int) (n int) {
	defer func(i int) {
		n = n + i
	}(i)
	
	i = i * 2
	n = i
	
	return
}

func testInterface(){
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	
	hi := "hi"
	fmt.Println(peo.Speak(hi))
}
