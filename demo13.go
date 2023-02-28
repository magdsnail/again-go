package main

import "fmt"

func main() {
	a := 1
	b := 5.0

	c := float64(a)
	d := int(b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n %d", d, d)
}
