package main

import (
	"github.com/pkg/errors"
	"log"
)

type Status struct {
	ok bool
}

func main() {
	postHandler()
}

// 控制器层
func postHandler() Status {
	err := insert("test")
	if err != nil {
		out := errors.Wrap(err, "I am the out error")
		log.Print(out)
		// 使用 +v 可以打印出调用栈
		//log.Printf("%+v", out)
		return Status{ok: false}
	}
	return Status{ok: true}
}

// service 层
func insert(contract string) error {
	err := dbQuery()
	if err != nil {
		// 也可以自定义参数
		return errors.Wrapf(err, "I am the middle error")
	}
	return nil
}

// model 层
func dbQuery() error {
	// Do something then fail
	err := errors.New("get db conn err")
	log.Println("new", err)
	return errors.Wrap(err, "I am inner error")
}
