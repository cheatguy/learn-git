package main

import (
	"fmt"
)
var channel =make (chan int ,100)  //全局变量
var (
myres = make(map[int]int, 20)
//close the lock
)

func factorial(n int) {
var res = 1
for i := 1; i <= n; i++ {
res *= i
}
channel <-res
//diguihanshu
myres [n] =<- channel
}

func main() {


for i := 1; i <= 20; i++ {
go factorial(i)
}

for i, v := range myres {
fmt.Printf("myres[%d] = %d\n", i, v)
}

}