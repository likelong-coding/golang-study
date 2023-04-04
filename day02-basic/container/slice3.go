package main

import "fmt"

func main() {

	var student = []string{"Tom", "Jack", "Lucy", "Likelong", "Mary"}

	for i, v := range student {
		fmt.Println(i, "->", v)
	}

	fmt.Println(student[3:])

	// 删除Lucy
	student = append(student[0:2], student[3:]...)

	fmt.Println(student)
	fmt.Println("student切片长度：", len(student))
	fmt.Println("student切片容量：", cap(student))

	var ids []int

	ids = append(ids, 1)
	ids = append(ids, 1)
	ids = append(ids, 1)
	fmt.Println(ids)
}
