package main

import "fmt"

func main() {
	var a bool = true
	var b bool = true

	if a == true && b == true {
		fmt.Println(a && b)
	}
	c := [5]int{2, 3}
	fmt.Println(c)
	fmt.Println(a || b)
}
