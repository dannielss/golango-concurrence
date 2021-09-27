package main

import (
	"fmt"
	"time"
)

func count1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Count 1:", i)
	}
}

func count2() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Count 2:", i)
	}
}

func main() {
	go count1()
	count2()
}
