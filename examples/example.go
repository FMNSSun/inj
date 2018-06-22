package main

import (
	"fmt"
	. "github.com/FMNSSun/inj"
)

type Point struct {
	X int `inject:"Point.x"`
	Y int `inject:"Point.y"`
	Shape Rectangle `inject:"Rectangle"`
}

type Rectangle struct {
	W int
	H int
}

func main() {
	test := true

	inj := NewInjector()

	inj.Register("Rectangle", func()interface{}{return Rectangle{}})

	if test {
		inj.Register("Point.x", func()interface{}{return 10})
		inj.Register("Point.y", func()interface{}{return 10})
	} else {
		inj.Register("Point.x", func()interface{}{return 222})
		inj.Register("Point.y", func()interface{}{return 222})		
	}

	p := Point{}
	inj.Inject(&p)
	fmt.Println(p)
}
