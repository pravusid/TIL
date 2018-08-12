package main

import (
	"fmt"
	"time"
)

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

func main() {
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("별도로 실행됨!")
	})
	fmt.Println("Hello World")
	CountDown(5)
}
