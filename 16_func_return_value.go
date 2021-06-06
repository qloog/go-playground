package main

import "fmt"

func List() (int, int, error) {

	list := 10
	total := 10

	return list, total, nil
}

func main() {

	list, total, _ := List()

	fmt.Println(list, total)

}
