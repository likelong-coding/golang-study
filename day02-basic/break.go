package main

import "fmt"

func main() {
	i := 1

OuterLoop:

	for {
		for {
			if i > 5 {
				// 跳出OuterLoop标签对应的循环，相当于直接跳出最外层循环
				break OuterLoop
			}
			fmt.Println(i)
			i++
		}
	}
}
