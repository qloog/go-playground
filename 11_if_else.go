package main

import "fmt"

func GetUserId() int {
	return 0
}

func main() {

	score := 90

	if score < 60 {
		fmt.Println("bad")
	} else if score >= 60 && score <= 80 {
		fmt.Println("good")
	} else {
		fmt.Println("very good")
	}

	if userId := GetUserId(); userId == 0 {
		fmt.Println("userId", userId)
	}

}
