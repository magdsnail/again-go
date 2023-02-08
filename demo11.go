package main

import "fmt"

func main() {
	var age int = 18

	//默认打印6位小数 尽量使用float64 来定义float类型
	var money float64 = 3.14

	fmt.Printf("%T, %d\n", age, age)

	//fmt.Printf("%T, %f\n", money, money)
	fmt.Printf("%T, %.2f\n", money, money)

}
