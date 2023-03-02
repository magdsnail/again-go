package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64

	//fmt.Println()
	//fmt.Printf()
	//fmt.Print()

	fmt.Println("请输入两个数1.整数 2.浮点数")

	//变量取地址 &变量
	fmt.Scanln(&x, &y)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	//fmt.Scanln() //接收输入 Scan |
}
