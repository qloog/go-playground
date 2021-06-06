package main

import "fmt"

type UserInfo struct {
	Id      int
	Name    string
	Age     int
	Country string
}

func main() {

	// 1、声明
	var userMap map[int]UserInfo
	fmt.Println("userMap", userMap)

	// 2、make 创建
	userMap1 := make(map[int]UserInfo, 1)
	fmt.Println("userMap1", userMap1)

	// 3. 创建并初始化
	userMap2 := map[int]UserInfo{
		1: UserInfo{1, "a", 20, "CN"},
		2: UserInfo{2, "b", 20, "CN"},
	}
	fmt.Println("userMap2", userMap2)

	// 4. 元素赋值
	userMap2[3] = UserInfo{3, "c", 21, "CN"}
	userMap2[4] = UserInfo{Id: 4, Name: "d", Country: "CN"}
	fmt.Println("userMap2", userMap2)

	// 5. 删除元素
	delete(userMap2, 2)
	fmt.Println("userMap2", userMap2)

	// 6. 元素查找
	val, ok := userMap2[5]
	if ok {
		fmt.Println("val", val)
	}

}
