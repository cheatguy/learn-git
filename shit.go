package main

import (
	"fmt"
	"time"
)

func main(){

	var channel = make(chan string, 3)

	go func() {
		channel <- "A"
		channel <- "B"
		channel <- "C"
		fmt.Println("我发送了3个数据")

		channel <- "D"
		fmt.Println("我发送了第4个数据")
	}()
	time.Sleep(time.Second)
}