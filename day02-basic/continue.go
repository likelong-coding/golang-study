package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	//fmt.Println(start)
OuterLoop:
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {

			if j == 1 {
				println(i, j)
				continue OuterLoop
			}
		}
	}

	end := time.Now()

	fmt.Printf("执行耗时：%v\n", end.Sub(start))
}
