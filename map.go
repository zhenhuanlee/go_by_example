package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := map[string]int{"a": 1, "b": 2}

	m1["a"] = 1
	delete(m2, "a")

	fmt.Println(m1)
	fmt.Println(m2)
}
