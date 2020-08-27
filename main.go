package main

import (
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

func main(){
	var err = cmd.Execute()
	if err!=nil{
		log.Fatalf("cmd.Execute err: %v",err)
}
}