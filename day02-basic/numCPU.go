package main

import (
	"fmt"
	"runtime"
)

func main() {

	// runtime包，使用NumCPU()函数获取了程序使用的CPU核数
	if num := runtime.NumCPU(); num >= 1 {
		fmt.Println("程序使用的CPU核数为：", num)
	}
}
