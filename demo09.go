package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "hahaha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c)
}
