package main

import "fmt"

func main() {
	// 二进制 0 1 逢二进一

	//电路问题

	//位运算 二进制上的 0 false 1 true
	// & 我和你 1 1 结果才为1     0,1
	// | 我或者你 0 0 结果才为1
	// 60 0011 1100
	// 13 0000 1101
	//------------------
	// &  0000 1100  我和你 同时满足
	// |  0011 1101  我或你 一个满足即可
	// ^  0011 0001  不同为1相同为0
	// >> 2
	// << 2

	var a uint = 60
	var b uint = 13

	var c uint = 0

	c = a & b
	fmt.Printf("%d, 二进制%b", c, c)

	c = a | b
	fmt.Printf("%d, 二进制%b", c, c)

	c = a ^ b
	fmt.Printf("%d, 二进制%b", c, c)

	// 60 0011 1100
	c = a << 2
	fmt.Printf("%d, 二进制%b", c, c) //11110000

	a = 60
	c = a >> 2
	fmt.Printf("%d, 二进制%b", c, c) //0000 1111
}
