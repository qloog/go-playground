package main

import (
	"fmt"
)

func add(args ...int) (sum int) {

	for _, arg := range args {
		sum = sum + arg
	}

	return sum
}

func main() {

	d := add(1, 2, 3, 4, 5)
	fmt.Println("d value:", d, "test")

}
