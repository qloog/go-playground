package main

import "fmt"

func main() {

	//1
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	// 2
	j := 0
	for j <= 3 {
		fmt.Println(j)
		j++
	}

	// do while, while(true) 死循环
	for {
		fmt.Println("loop")
		break
	}

	sum := 0
	for {
		sum++
		if sum >= 10 {
			fmt.Println(sum)
			break
		}
	}

	//无限双循环

JumpLoop:

	for {
		for {
			fmt.Println("double loop")
			break JumpLoop
		}
	}

JLoop:
	for i := 0; i <= 6; i++ {
		for j := 0; j <= 5; j++ {
			if i > 4 {
				break JLoop
			}
			fmt.Println(i)
		}
	}

	// array, slice, map
	// for + range
	demoArr := [5]int{1, 2, 3, 4, 5}
	for index, value := range demoArr {
		fmt.Println(index, value)
	}

	demoSlice := []int{6, 7, 8}
	for index, value := range demoSlice {
		fmt.Println(index, value)
	}

	demoMap := make(map[string]int, 5)
	demoMap["a"] = 1
	demoMap["b"] = 2
	for index, value := range demoMap {
		fmt.Println(index, value)
	}

}
