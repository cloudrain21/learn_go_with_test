package main

import (
	"fmt"
	"reflect"
)

func f(a, b interface{}) (concat string) {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)

	switch v1.Kind() {
	case reflect.Int:
		x := v1.Int() + v2.Int()
		concat = fmt.Sprintf("%v", x)
	case reflect.String:
		concat = v1.String() + v2.String()
	}
	return concat
}

func main() {
	r := f(1, 2)
	fmt.Printf("%v\n", r)

	r = f("1", "2")
	fmt.Printf("%v\n", r)
}
