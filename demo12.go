package main

import "fmt"

func main() {
	var str string
	str = "hello, kuangshen"
	fmt.Printf("%T, %s\n", str, str)

	//单引号

	//编码表 ASCII字符码
	//扩展
	//所有的中国字的编码表 GBK
	//全世界的编码表：Unicode编码表
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T, %d\n", v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	//字符串连接
	fmt.Println("hello" + ",xuexiangban")

	//转义字符 \
	fmt.Println("hello\"xuexiangban")
	fmt.Println("hello\nxuexiangban") //n换行
	fmt.Println("hello\txuexiangban") //t tab制表符
}
