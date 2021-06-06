package main

import (
	"fmt"
)

func main() {
	// 定义
	var isVip bool

	isVip = true

	fmt.Println("isVip is:", isVip)

	// 不支持自动转换
	//isVip = 1
	//fmt.Println("isVip is:", isVip)

	//isVip = bool(1)
	//fmt.Println("isVip is:", isVip)
}
