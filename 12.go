package main

import (
	"fmt"
)

func main() {
	//12 задача
	var a int
	fmt.Scan(&a)
	for i := a; i > 0; i -= 1 {
		fmt.Println(i)
	}
}
