package main

import (
	"fmt"
	. "github.com/FMNSSun/inj"
)

func loop() {
	i := CreateOrNil("count").(int)
	for j := 0; j < i; j++ {
		fmt.Println(j)
	}
}

func main() {
	inj := NewInjector()
	inj.Register("count", func()interface{}{return 10})
	SetInjector(inj)
	loop()
}
