package main

import "fmt"

func main() {

	var i = 10

Loop:

	for i < 20 {
		if i == 16 {
			i = i + 1
			fmt.Println("---")
			goto Loop
		}

		fmt.Println(i)
		i++
	}
}
