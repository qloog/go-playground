package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x int32
	x = 1
	//

	//x := 1

	fmt.Println("x is:", x)

	// + , 1 * / %
	// int, int32

	var y int32
	y = 2

	z := x + y
	fmt.Println("z is:", z)
	fmt.Println("z type:", reflect.TypeOf(z))

	var i int32
	var j int64

	i = 1
	j = 2

	if j == i {
		fmt.Println("equal")
	}
}
