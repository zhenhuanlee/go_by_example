package main

import "fmt"

var vec [5]int = [5]int{2, 3, 12, 0, 5}

func main() {
	for i, foo := range vec {
		for j, too := range vec[i:] {
			if too > foo {
				vec[j], vec[i] = vec[i], vec[j]
			}
			fmt.Println(vec)
		}
	}
}