package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{}, 3)
	go chDemo(c)
	for i := 0; i < 3; i++ {
		c <- struct{}{}
		time.Sleep(2 * time.Second)
	}
	// ctx, cacel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cacel()
	// go func() {
	// 	for {

	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("--------Done----------")
	// 			fmt.Println("ctx error: ", ctx.Err())
	// 		default:
	// 		}
	// 		time.Sleep(200 * time.Millisecond)
	// 		fmt.Println("running....")
	// 	}
	// }()
	// time.Sleep(3 * time.Second)
	// panic("running err...")
}

func chDemo(c chan struct{}) {
	for {
		t := time.After(3 * time.Second)
		select {
		case <-c:
			fmt.Println("running....")
		case <-t:
			fmt.Println("timeout...")
			time.Sleep(200 * time.Microsecond)
		default:
		}
	}
}
