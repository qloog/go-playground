package main

import "fmt"

func main() {

	// 1. 定义 slice
	var intSlice []int
	fmt.Println("slice:", intSlice)

	// 2. 创建切片
	// a. 直接创建
	slice1 := make([]int, 5, 5)
	fmt.Println("slice1:", slice1)

	slice2 := make([]int, 5)
	fmt.Println("slice2:", slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice3:", slice3)

	// b. 基于数组创建
	intArr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	intSlice2 := intArr[:5]
	fmt.Println("intslice2", intSlice2)
	fmt.Printf("slice2 type %T", intSlice2)

	// c. 基于切片创建
	slice4 := slice3[:3]
	fmt.Println("slice4:", slice4)

	// 3. 遍历切片元素
	// a. for
	for i := 0; i < len(slice4); i++ {
		fmt.Println("value :", slice4[i])
	}

	// b. range
	for index, value := range slice4 {
		fmt.Printf("index: %d, value %d", index, value)
	}

	// 查看切片的长度和容量
	fmt.Println("slice4 len", len(slice4))
	fmt.Println("slice4 cap", cap(slice4))

	// 4. 动态增减元素
	// append
	slice4 = append(slice4, 4, 5, 6, 7, 8)
	fmt.Println("slice4", slice4)

	slice3 = append(slice3, slice2...) // 打散切片
	fmt.Println("slice3", slice3)

	// 5. 内容复制

	newSlice1 := []int{11, 12, 13, 14, 15}
	newSlice2 := []int{21, 22, 23}
	//copy(newSlice1, newSlice2)
	//fmt.Println("newSlice1", newSlice1)

	copy(newSlice2, newSlice1)
	fmt.Println("newSlice2", newSlice2)
}
