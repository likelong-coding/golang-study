package main

import "fmt"

func main() {
	str := "Java语言"

	bytes := []byte(str)
	for i := 0; i < 4; i++ {
		bytes[i] = ' '
	}

	fmt.Println(string(bytes))
}
