package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	timeout := time.Second * 2
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	done := make(chan int, 1)

	// do work
	go func() {
		fmt.Println("do working")
		//time.Sleep(time.Second)
		// 超时
		time.Sleep(time.Second * 3)

		// work done
		done <- 1
	}()

	// 阻塞，等待
	select {
	// 统一超时处理
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-done:
		fmt.Println("work done on time")
	}

	fmt.Println("main func exit")
}
