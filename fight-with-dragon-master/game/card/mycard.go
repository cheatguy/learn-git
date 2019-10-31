// Time : 2019/10/19 20:28
// Author : MashiroC

// card
package card

import (
	"fmt"
	"mashiroc.fun/dragongame/game"
	"mashiroc.fun/dragongame/game/effect"
)

// mycard.go something

type WildBoar struct {  //这是一个随从
	game.BaseCard
}

func NewWildBoar() game.Card {   //这是实现他的方法，但是方法名字不能重名
	return &WildBoar{
		*game.NewCard("石牙野猪", 1, 1, 1, effect.Charge{}),
	}
}

//type WitchDoctor struct{
//	game.BaseCard
//}
//
//func WitchDoctor(){
//
//}
type WolfCavalry struct{
	game.BaseCard
}

func WolfCavalrykik() game.Card {
	return &WolfCavalry {
		*game.WolfCavalry("狼骑兵", 3, 3, 1, effect.Charge{}),
	}
}


type ElvenArcher struct {
	game.BaseCard
}

func NewElvenArcher() game.Card {
	return &ElvenArcher{
		*game.NewCard("精灵弓箭手",
			1,
			1,
			1,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个对方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = dragon
					} else {
						target = dragon.Servant()[flag-1]
					}
					target.ReduceHp(1)
				}),
		),
	}
}
