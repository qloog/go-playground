package main

import (
	"fmt"
)

var (
	v1 int
	v2 string
)

func main() {

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	// 变量初始化
	var v3 int = 3
	fmt.Println("v3:", v3)

	var v4 = 4
	fmt.Println("v4:", v4)

	v5 := 5
	fmt.Println("v5:", v5)

	// 变量赋值

	var v6 int
	v6 = 6
	fmt.Println("v6:", v6)

	// 注意
	//var v7 int
	//v7 := 7
	//fmt.Println("v7:", v7)

	var v8, v9 int = 8, 9
	fmt.Println(v8, v9)

	var v10, v11 = 10, 11
	fmt.Println(v10, v11)

	v12, v13 := 12, 13
	fmt.Println(v12, v13)

	var v14, v15 string = "v14", "v15"
	fmt.Println(v14, v15)

}
