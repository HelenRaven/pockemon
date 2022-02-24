package pockemon

import (
	"fmt"

	_ "github.com/valyala/fasthttp"
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
