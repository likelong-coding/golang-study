package main

import "fmt"

func main() {

	var name string
	// 同时给多个变量赋值
	name, age := "Tom", 18

	fmt.Println("name: ", name)
	fmt.Println("age: ", age)

}
