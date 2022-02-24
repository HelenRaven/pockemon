package pockemon

import (
	"fmt"
)

type Pokemon struct {
	Name string
}

func (p *Pokemon) IChooseU() {
	fmt.Println("I choose u,", p.Name)
}
