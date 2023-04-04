package main

import (
	"fmt"
	"reflect"
)

func main() {

	a, b := 3.0, 2.0
	fmt.Println(a / b)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

	str := "hello, 你好"
	fmt.Println(str)

	line := `第一行
第二行
第三行
`
	fmt.Print(line)
}
