package main

import "fmt"

func main() {
	x := make(map[string]int, 10)

	for i := 0; i < 10; i++ {
		x[fmt.Sprintf("key-%d", i)] = i
	}

	for key, value := range x {
		fmt.Println(key, value)
	}
}
