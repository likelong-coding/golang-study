package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	a, b := "012345", "6789"

	// 字符串拼接
	fmt.Println(a + b)

	// 当需要拼接的字符串较长时，
	//使用“+”操作符进行字符串的拼接并不高效，因此推荐使用字节缓冲的方式来进行。

	// 声明变量c，类型为字节缓冲
	var c bytes.Buffer
	// 写入字符串变量a内容
	c.WriteString(a)
	// 写入字符串变量b内容
	c.WriteString(b)

	fmt.Println(c.String())
	fmt.Println(reflect.TypeOf(c))
}
