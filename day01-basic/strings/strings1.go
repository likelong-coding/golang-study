package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Go语言,Python语言"

	index := strings.LastIndex(str, "语")
	fmt.Println(index)

	fmt.Println(str[index:])
}
