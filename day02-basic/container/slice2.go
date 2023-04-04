package main

import "fmt"

func main() {

	var student = []string{"Tom", "Jack"}

	fmt.Println("student切片：", student)
	fmt.Println("切片长度", len(student))
	fmt.Println("切片容量", cap(student))
	fmt.Println("判断切片是否为空：", student == nil)
}
