package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	v map[string]int
}

func main() {
	var t Test
	t.v = make(map[string]int)

	fmt.Println(unsafe.Sizeof(t))
	t.v["a"]++

	t2 := map[string]int{}
	fmt.Println(unsafe.Sizeof(t2))
	t2["a"]++
	fmt.Println(t2, unsafe.Sizeof(t2))
	t2["b"] = 0
	fmt.Println(t2, unsafe.Sizeof(t2))

	var arr []string
	fmt.Println(unsafe.Sizeof(arr))
	arr = append(arr, "aaa")
	fmt.Printf("arr : %v\n", arr)

	arr1 := make([]string, 1)
	arr1 = append(arr1, "xxx")
	fmt.Printf("arr : %v\n", arr1)
	arr1 = append(arr1, "yyy")
	fmt.Printf("arr : %v\n", arr1)
}
