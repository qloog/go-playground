package main

import (
	"github.com/dropbox/godropbox/errors"
	"log"
)

func main() {
	replyHandler()
}

// 控制器层
func replyHandler() {
	err := replyLogic("test")
	if err != nil {
		out := errors.Wrap(err, "I am the out error")
		log.Println(out.Error())
		//log.Println(out.GetStack())
	}
}

// service 层
func replyLogic(contract string) error {
	err := create()
	if err != nil {
		// 也可以自定义参数
		return errors.Wrap(err, "I am the middle error")
	}
	return nil
}

// model 层
func create() error {
	// Do something then fail
	return errors.New("I am inner error")
}
