package main

import (
	"fmt"
	"time"
)

func main() {

	orderStatus := 2

	switch orderStatus {
	case 0:
		fmt.Println("padding")
	case 1:
		fmt.Println("paid")
		fallthrough
	case 2:
		fmt.Println("complete")
	default:
		fmt.Println("exception")

	}

	//t := time.Now()
	switch t := time.Now(); {
	case t.Hour() > 12:
		fmt.Println("after noon")
	default:
		fmt.Println("befer noon")
	}

}
