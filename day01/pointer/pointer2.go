package main

import "fmt"

func main() {
	var p *int
	p = new(int)

	fmt.Println("指针p所指向的内容:", *p)

	*p = 10
	fmt.Println("指针p所指向的内容:", *p)

}
