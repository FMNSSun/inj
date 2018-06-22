package main

import (
	"fmt"
	. "github.com/FMNSSun/inj"
)

type Owner interface {
	Name() string
}

type ConcreteOwner struct {
}

func (_ *ConcreteOwner) Name() string { return "ConcreteOwner" }

type Animal interface {
	Sound() string
}

type Dog struct {
	CO Owner `inject:"Owner"`
}

func (d *Dog) Sound() string { return d.CO.Name() + ": Woof!" }

type Cat struct {
	CO Owner `inject:"Owner"`
}

func (c *Cat) Sound() string { return c.CO.Name() + ": Meow!" }

type A struct {
	SomeAnimal Animal `inject:"Animal"`
	a int
}

func main() {
	inj := NewInjector()
	inj.Register("Animal", func()interface{} { return &Dog{} })
	inj.Register("Owner", func()interface{} { return &ConcreteOwner{} })
	inj.Register("A", func()interface{} { return &A{} })
	SetInjector(inj)
	
	a := inj.CreateOrNil("A").(*A)

	fmt.Println(a.SomeAnimal.Sound())
}
