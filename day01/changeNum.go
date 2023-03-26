package main

import "fmt"

func main() {

	// 交换变量值

	//var c int
	a, b := 2, 3

	//c = a
	//a = b
	//b = c

	a, b = b, a

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

}
