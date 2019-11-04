package main

import (
	"Crontab/crontabTask"
	"fmt"
)
//延迟消息

func test1(params ...interface{}){
	fmt.Println("test1: ", params)
}

func main() {
	crontabTask.CheckIfConnect("http://127.0.0.1:8500")
	select {

	}
}