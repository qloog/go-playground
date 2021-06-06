package main

import (
	"fmt"
	"sync"
)

func main() {
	// 互斥锁保护计数器
	var mu sync.Mutex
	// 计数器的值
	var count = 0
	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				// 只要是多线程或者多goroutine对同一共享资源(counter)进行访问，就会存在 数据竞争 问题
				// 最简单的处理方式就是加锁，虽然会有一定的性能损耗，但是通过损失性能保证数据的正确性还是很有必要的
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
