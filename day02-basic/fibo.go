package main

import "fmt"

func main() {

	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci1(5))
}

/*
斐波那契数列递归实现
*/
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/*
简单动态规划
*/
func fibonacci1(n int) int {

	a := 1
	b := 1

	for i := 2; i < n; i++ {
		c := b
		b = a + b
		a = c
	}
	return b
}
