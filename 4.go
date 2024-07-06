package main

import (
	"fmt"
)

func three(a, b, c int) int {
	return max(a, b, c)
}

func main() {
	//4 задача
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(three(x, y, z))
}
