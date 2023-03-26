package main

import "fmt"

func main() {

	// 匿名变量

	a, _ := returnData()
	_, b := returnData()
	fmt.Print(a, b)
}

func returnData() (int, int) {
	return 20, 10
}
