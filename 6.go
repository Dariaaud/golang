package main

import (
	"fmt"
)

func vov(c string) string {
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" || c == "y" {
		return ("Yes")
	} else {
		return ("No")
	}
}

func main() {
	//6 задача
	var c string
	fmt.Scan(&c)
	fmt.Println(vov(c))
}
