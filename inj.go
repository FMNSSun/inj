package inj

import (
	"fmt"
	"reflect"
)

var injector *Injector

// Register an Injector as the global Injector
func SetInjector(inj *Injector) {
	injector = inj
}

// Invokes the global Injector's Inject method.
func Inject(a interface{}) error {
	return injector.Inject(a)
}

// An Injector allows one to `inject` values
// into a struct.
type Injector struct {
	injects map[string]func()interface{}
}

// Register an injection function under the specified name.
// This must match the struct tags.
func (f *Injector) Register(name string, fn func()interface{}) {
	f.injects[name] = fn
}

// Inject the values into the supplied struct.
func (f *Injector) Inject(a interface{}) (err error) {
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("Error in Inject: %v", err.Error())
		}
	}()

	f.inject(a)
	
	return
}

// Invokes the global Injector's CreateOrNil function
func CreateOrNil(name string) interface{} {
	return injector.CreateOrNil(name)
}

// Like create, but returns nil to indicate error.
func (f *Injector) CreateOrNil(name string) interface{} {
	it, err := f.Create(name)

	if err != nil {
		return nil
	}

	return it
}

// Invokes the global Injector's Create function. 
func Create(name string) (interface{}, error) {
	return injector.Create(name)
}

// Creates using the function registered under the specified name. 
func (f *Injector) Create(name string) (interface{}, error) {
	fn, ok := f.injects[name]

	if !ok {
		return nil, fmt.Errorf("%v is not registered!", name)
	}

	it := fn()

	err := f.Inject(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// Returns a new Injector
func NewInjector() *Injector { 
	return &Injector{injects: make(map[string]func()interface{})} 
}

func (f *Injector) inject(a interface{}) {
	var rt reflect.Type
	var rv reflect.Value

	rt = reflect.TypeOf(a)
	rv = reflect.ValueOf(a)

	if rt.Kind() == reflect.Ptr {
		rt = reflect.TypeOf(a).Elem()
		rv = reflect.ValueOf(a).Elem()
	}

	if rt.Kind() != reflect.Struct {
		return
	}
	
	nf := rt.NumField()
	
	for i := 0; i < nf; i++ {
		fld := rt.Field(i)

		inj, _ := fld.Tag.Lookup("inject")
		
		if inj == "" {
			continue
		}
		
		fn, ok := f.injects[inj]
		
		if !ok || fn == nil {
			panic(fmt.Sprintf("%v is not registired", inj))
		}
		
		it := fn()
		
		it_rv := reflect.ValueOf(it)

		rv.Field(i).Set(it_rv)


		if it_rv.Kind() == reflect.Ptr || it_rv.Kind() == reflect.Struct {
			f.inject(it)
		}
	}
}
