package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.4

	// value -> type
	fmt.Println("reflect.Type", reflect.TypeOf(x)) // float64

	// value -> value wrapper
	v := reflect.ValueOf(x)

	fmt.Println("Value", v)
	fmt.Println("Value to String", v.String())
	fmt.Println("Value to Type", v.Type())
	fmt.Println("Value to Kind", v.Kind() == reflect.Float64)
	fmt.Println("Value to Float", v.Float())

	// value wrapper -> value
	y := v.Interface().(float64)
	fmt.Println(y) // 3.4

	access()
}

func access() {
	var x float64 = 3.4
	// it creates a copy but you need a link
	v := reflect.ValueOf(x)
	// v.SetFloat(4.3) // Error panic
	fmt.Println("Can Set?", v.CanSet()) // false

	// send x pointer to get access by link
	p := reflect.ValueOf(&x)
	fmt.Println("pointer type", p.Type())
	fmt.Println("Can Set?", p.CanSet()) // false

	// than take a link to get real value
	e := p.Elem()
	fmt.Println("Can Set?", e.CanSet()) // true
	e.SetFloat(4.3)
	fmt.Println(x) // 4.3
}
