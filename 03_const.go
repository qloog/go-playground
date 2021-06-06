package main

import (
	"fmt"
)

func main() {

	const age = 0

	fmt.Println("age:", age)

	//const uid = os.Getgid()
	//fmt.Println("uid:", uid)

	const mask = 1 << 3
	fmt.Println("mask:", mask)

	// 分组
	const (
		height int64 = 100
		weight int64 = 100
	)

	// 预定义常量
	// true false iota

	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Println(c0, c1, c2)

	const x = iota
	const y = iota
	fmt.Println(x, y)

	const (
		c3 = iota
		c4
		c5
	)
	fmt.Println(c3, c4, c5)

	const (
		Sunday = iota
		Monday
		Tuesday
	)
	fmt.Println(Sunday, Monday, Tuesday)

}
