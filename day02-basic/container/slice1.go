package main

import "fmt"

func main() {

	student := [...]string{"Tom", "Ben", "Peter"}

	student1 := student[0:1]

	fmt.Println("student数组：", student)
	fmt.Println("student1切片：", student1)
	fmt.Println(student1[:])
	fmt.Println("student数组地址为", &student[0])   //取student[0]元素的地址
	fmt.Println("student1切片地址为", &student1[0]) //取student1[0]元素的地址
	fmt.Println("student1切片长度为：", len(student1))
	fmt.Println("student1切片容量为：", cap(student1))

}
