package main

import (
	"fmt"
	"reflect"
)

func main() {
	proc := func(p interface{}) {
		switch t := p.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("unknown", reflect.TypeOf(t))
		}
	}

	proc(1)
	proc("ha")
	proc(true)
}
