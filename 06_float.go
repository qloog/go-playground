package main

import (
	"fmt"
	"math"
	"reflect"
)

// p 比较精度 0.00001
func IsEqual(f1, f2, p float64) bool {
	return math.Abs(f1-f2) < p
}

func main() {

	var f1 float32
	f1 = 1.23

	f2 := 1.234

	fmt.Println("f1:", f1)
	fmt.Println("f2:", f2)

	f1 = float32(f2)

	fmt.Println("f2 type:", reflect.TypeOf(f2))
	fmt.Println("f1:", f1)

	// > < >= <=
	// ==
	a := 1.2345678
	b := 1.2345670
	// 0.00000008
	p := 0.00001

	fmt.Println("is equal:", IsEqual(a, b, p))

}
