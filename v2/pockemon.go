package pockemon

import (
	"fmt"
	"math/rand"

	_ "github.com/gorilla/websocket"
)

type Pokemon struct {
	Name string
}

func (p *Pokemon) IChooseU() {
	fmt.Println("I choose u,", p.Name)
}

func StartFight(p1, p2 *Pokemon) {
	fmt.Printf("Start fight! %s VS %s", p1.Name, p2.Name)
}

func Winner(p1, p2 *Pokemon) {
	res := rand.Intn(3)

	switch res {
	case 0:
		fmt.Println("Draw game")
	case 1:
		fmt.Println("Winner:", p1.Name)
	case 2:
		fmt.Println("Winner:", p2.Name)
	}
}
