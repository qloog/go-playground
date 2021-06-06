package main

import (
	"fmt"
	"time"
)

// link: https://segmentfault.com/a/1190000015084958
func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)

		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
		//default:
		//	fmt.Println("default")
	}
}
