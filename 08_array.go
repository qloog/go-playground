package main

import "fmt"

func main() {

	// 1. 定义
	var arr1 [6]int
	fmt.Println("arr1", arr1)

	// 2. 定义及初始化
	arr2 := [5]int{5, 6, 7, 8, 9}
	fmt.Println("arr2", arr2)

	// 3. 访问 0, len(array)-1
	fmt.Println("arr2 1:", arr2[1])

	// 4. 修改
	arr2[2] = 700
	fmt.Println("arr2 2:", arr2[2])
	fmt.Println("arr2", arr2)

	// 5. 遍历
	// a. for
	fmt.Println("for arr2...")
	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}
	// b. range
	fmt.Println("range arr2...")
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

}
