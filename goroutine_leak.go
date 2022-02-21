package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

type People struct {
	Name string `json:"name"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:6060",
	}

	http.HandleFunc("/", Bad)
	http.HandleFunc("/fine", Good)

	_ = server.ListenAndServe()

}

func Bad(w http.ResponseWriter, r *http.Request) {
	// 定义一个大小的chan, 实际写入了3个，会有2个泄露
	// 通过使用 http://127.0.0.1:6060/debug/pprof/ 可以进行查看
	errChan := make(chan error, 1)
	finishChan := make(chan struct{}, 1)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		errChan <- fmt.Errorf("test01")
	}()

	go func() {
		defer wg.Done()
		errChan <- fmt.Errorf("test02")
	}()

	go func() {
		defer wg.Done()
		errChan <- fmt.Errorf("test03")
	}()

	go func() {
		wg.Wait()
		//close(errChan)
		close(finishChan)
	}()

	//for {
	select {
	case <-finishChan:
	case err := <-errChan:
		fmt.Println("err", err.Error())
	}
	//}

	_, _ = fmt.Fprintf(w, "Hello world!!!")
}

func Good(w http.ResponseWriter, r *http.Request) {
	// 正确的案例
	errChan := make(chan error, 1)
	finishChan := make(chan struct{}, 1)

	go func() {
		// 必须使用for 遍历errChan， 否则就会泄露
		for {
			select {
			case <-finishChan:
				return
			case err := <-errChan:
				if err != nil {
					fmt.Println("err", err.Error())
				}
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("go 1 leaving...")
		}()
		errChan <- fmt.Errorf("test01")
	}()

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("go 2 leaving...")
		}()
		errChan <- fmt.Errorf("test02")
	}()

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("go 3 leaving...")
		}()
		errChan <- fmt.Errorf("test03")
	}()

	wg.Wait()
	close(errChan)
	close(finishChan)

	_, _ = fmt.Fprintf(w, "Hello world!!!")
}
