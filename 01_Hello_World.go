package main

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Println("Hello World!")

	a := fmt.Sprintf("%03d", 1)
	fmt.Println(a)

	fmt.Println(time.Now().Day())

	if 1547619571 >= 1547568000 && 1547619571 <= 1547598600 {
		fmt.Println("ok")
	} else {
		fmt.Println("no ok")
	}
}
