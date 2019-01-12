package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	for i, val := range nums {
		fmt.Println(i, "-", val)
	}
	fmt.Println("------------")

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, "-", v)
	}
	fmt.Println("------------")

	for k := range m {
		fmt.Println(k)
	}
	fmt.Println("------------")

	for i, c := range "go" {
		fmt.Println(i, "--", c)
	}
}
