// Time : 2019/10/19 15:47
// Author : MashiroC

// character
package game

import (
	"fmt"
	"mashiroc.fun/dragongame/color"
)

// character.go something

const (
	maxServant = 3
)

type Character interface {  //角色接口，或者是随从接口
	Card  //角色包含的方法
	Do(target Character)
	Servant() []Card  //召唤仆人
	CheckServant()
	Print()//
	AttendNum() int   //返回值为int
}

type baseCharacter struct {
	servant   []Card   //定义一个英雄结构体 比如猎人啊什么的
	name      string
	maxHp     int
	hp        int
	cost      int
	maxCost   int
	hurtNum   int
	attendNum int
}

func (c *baseCharacter) Begin() {     //传入进来一个随从
	c.maxCost++    //模拟每个回合开始
	c.cost = c.maxCost   //恢复费用
	c.attendNum = 1  //英雄自己也能进行攻击
}

func (c *baseCharacter) end() {
	//结束没有什么东西
}

func (c *baseCharacter) CheckServant() {
	//有一个迭代器
	for i, v := range c.servant {
		if v.Die() {//死了就删除这个东西
			c.servant = append(c.servant[:i], c.servant[i+1:]...)//他把随从放在一个切片里面
			return
		}
	}
}

func (c *baseCharacter) Attend(target Card) {   //经行攻击
	if c.attendNum == 0 {
		fmt.Println("你这回合不能进行攻击了。")
		return  //这是最初的判断
	}
	c.attendNum--//
	target.ReduceHp(c.hurtNum)//扣血
	c.ReduceHp(target.Hurt())
	return
}

func printCard(cards []Card, prefix string) {
	str := ""//？？？？应该是申明了一个空的字符串

	strName := fmt.Sprintf("%-8s\t", prefix)

	for i := 0; i < len(cards); i++ {
		strName += fmt.Sprintf("%-15s\t", cards[i].Name())  //是名字要加长吗
	}

	strName += "\n"
	str += strName

	strHp := "damage,cost,hp\t"
	for i := 0; i < len(cards); i++ {
		strHp += fmt.Sprintf("%s\t%s\t%s\t\t\t", color.Red(cards[i].Hurt()), color.Yellow(cards[i].Cost()), color.Green(cards[i].Hp()))
	}

	for i := len(cards); i < maxServant; i++ {
		strHp += "   \t\t\t"
	}
	strHp += "\n"
	str += strHp
	fmt.Print(str)
}

func (c *baseCharacter) Hp() int {
	return c.hp
}

func (c *baseCharacter) Name() string {
	return c.name
}

func (c *baseCharacter) Hurt() int {
	return c.hurtNum
}

func (c *baseCharacter) Cost() int {
	return c.cost
}

func (c *baseCharacter) ReduceHp(num int) {
	c.hp -= num
}

func (c *baseCharacter) Die() bool {
	return c.hp <= 0
}

func (c *baseCharacter) Servant() []Card {
	return c.servant
}

func (c *baseCharacter) AttendNum() int {
	return c.attendNum
}

func (c *baseCharacter) End() {
}

func (c *baseCharacter) SetAttendNum(n int) {
	c.attendNum = n
}

func (c *baseCharacter) BeforeDeploy() {
}
