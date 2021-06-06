package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 1. 声明 初始化
	var str string
	str = "Hello World"
	fmt.Println(str)

	// 2. 获取字符串
	str1 := str[0]
	fmt.Println(string(str1))

	// 对于字符的操作，操作的是字节数组

	// 3. 常用操作：连接， 长度， 获取
	// 连接
	a := "Hello"
	b := " Gopher"
	c := a + b
	fmt.Println(c)

	fmt.Println(len(a))

	// 4. 遍历
	// a. 字节数组的方式
	str2 := "hello 北京"
	for i := 0; i < len(str2); i++ {
		fmt.Println(string(str2[i]), reflect.TypeOf(str2[i]))
	}

	// b. 普通的遍历方式，unicode 字符遍历
	for index, data := range str2 {
		fmt.Println(index, string(data), reflect.TypeOf(data))
	}

	// 5. 修改字符串
	str3 := "hello Gopher"
	//str3[0] = "x"
	//fmt.Println(str3)

	b1 := []byte(str3)
	b1[0] = 'x'

	fmt.Println(string(b1))

}
