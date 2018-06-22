# Inj

Inj is a very minimal "injection framework". 

```go
type C1 struct {
	It *C2 `inject:"C2"`
}

type C2 struct {
	It *C3 `inject:"C3"`
}

type C3 struct {
	X int
}

func main() {
	inj := NewInjector()
	inj.Register("C2", func()interface{}{ return &C2{} })
	inj.Register("C3", func()interface{}{ return &C3{X: 10} })
	
	it := &C1{}
	inj.Inject(it)
	fmt.Println(it.It.It.X)
}
```
