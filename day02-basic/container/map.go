package main

import "fmt"

func main() {

	// 声明
	//var country map[string]string
	// make()函数初始化
	//country = make(map[string]string)

	country := make(map[string]string)
	country["中国"] = "China"
	country["日本"] = "Japan"

	fmt.Println(country)
	fmt.Println(len(country))

	//for _, v := range country {
	//	fmt.Println(v)
	//}

	delete(country, "日本")
	fmt.Println(country)
}
