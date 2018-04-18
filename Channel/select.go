package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "1 sec"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "2 sec"
	}()
	for i := 0; i < 2; i++ {
		select {
		case v1 := <-c1:
			fmt.Println(v1)
		case v2 := <-c2:
			fmt.Println(v2)
		}
	}

	//select {}  // block forever
	fmt.Println("start")
	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("do some work")
		}
	}()

	//select {}

	timeout := make(chan bool, 1)
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 2)
		timeout <- true
	}()
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-timeout:
		fmt.Println("timeout")
	}
}
