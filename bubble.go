package main

import "fmt"

var vec [5]int = [5]int{2, 3, 12, 0, 5}

func main() {
	var len = len(vec)

	for i:=0; i<len; i++ {
		for j:=i+1; j<len; j++ {
			fmt.Println(vec[i], "-", vec[j])
			if vec[i] > vec[j] {
				vec[j], vec[i] = vec[i], vec[j]
			}
			fmt.Println(vec)
		}
	}
}