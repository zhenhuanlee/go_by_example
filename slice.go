// key data type
package main

import "fmt"

func main() {
	var s = make([]string, 3)
	s[1] = "b"
	s = append(s, "a", "c")
	for _, s := range s {
		fmt.Println(s)
	}
}
