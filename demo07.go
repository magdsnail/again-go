package main

import "fmt"

func test() (int, int) {
	return 100, 200
}

func main() {
	//匿名变量 丢弃
	a, _ := test()
	_, b := test()
	fmt.Println(a)
	fmt.Println(b)
}
