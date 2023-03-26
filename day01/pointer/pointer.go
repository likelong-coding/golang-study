package main

import "fmt"

func main() {

	num := 1
	var p *int

	p = &num

	fmt.Println("num变量的地址值：", p)
	fmt.Println("指针变量p的地址值：", &p)

	fmt.Println("指针变量p所指向的内容:", *p)
}
