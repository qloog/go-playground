package main

import (
	"fmt"
	"time"
)

type People struct {
	Name string `json:"name"`
}

func main() {

	ret := time.Now().YearDay()
	fmt.Println("ret", ret)
}
