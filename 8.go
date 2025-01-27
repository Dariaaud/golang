package main

import (
	"fmt"
)

func reverse(c string) string {
	new_one := []rune(c)
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		new_one[i], new_one[j] = new_one[j], new_one[i]
	}
	return string(new_one)
}

func main() {
	//8 задача
	var c string
	fmt.Scan(&c)
	fmt.Println(reverse(c))
}
