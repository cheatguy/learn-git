package main

import (

	"fmt"
)

// person 人类
type person struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}


type lemon interface {
	acid()
}
// dove 鸽子
type dove interface {//jiekou
	gugugu() // 鸽只用声明一个方法就可以
}


type zhenxiang interface {
	xiang()
}

type repeater interface {
	repeat(string) // 复读
}

// repeat 复读
func (p person) repeat(word string) {
	fmt.Println(word)
}

func (p person) gugugu(){
	fmt.Println(p.name, "又鸽了")
}

func (p person) acid() {
	fmt.Println(p.name,"疯狂吃柠檬")
}
func (p person) xiang(){//让人类实现4个接口
	fmt.Println("哎呀，真香")
}
func main() {
     p:=person{
	name :"wjz",
	age: 19,
	gender: "male",
    }
     var le lemon
     var xi  zhenxiang
     var do dove
     var re  repeater
     //声明了4种接口
     le=p
     xi=p
     do=p
     re=p
     //对接口通过结构体转化进行初始化
     le.acid()
     xi.xiang()
     do.gugugu()
     re.repeat("deep dark fantasy")



}