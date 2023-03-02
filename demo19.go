package main

import "fmt"

func main() {
	//a := 0
	//for i := 1; i <= 10; i++ {
	//	a = a + i

	//}
	//fmt.Println(a)
	//for j := 0; j < 5; j++ {
	//	for i := 1; i <= 5; i++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d * %d = %d \t", i, j, i*9)
		}
		fmt.Println()
	}

}
