package main

import "fmt"

func main() {
	switch {
	// true，肯定执行
	case true:
		fmt.Println("case 1为true")
		fallthrough
	// 由于上一个case中有fallthrough，即使是false，也强制执行 case false:
	case false:
		fmt.Println("case 2为false")
		fallthrough
	default:
		fmt.Println("默认 case")
	}
}
