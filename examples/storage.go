package main

import (
	"fmt"
	. "github.com/FMNSSun/inj"
)

type Storage interface {
	Put(i int) 
	Get() int
}

type StorageA struct {
	v int
}

func (s *StorageA) Put(i int) {
	s.v = i
}

func (s *StorageA) Get() int {
	return s.v
}

func Do(i int) {
	st := CreateOrNil("Storage").(Storage)
	st.Put(i)
	fmt.Println(st.Get())
}

func main() {
	inj := NewInjector()
	inj.Register("Storage", func()interface{} { return &StorageA{} })
	SetInjector(inj)

	for i := 0; i < 3; i++ {
		Do(i)
	}
}
