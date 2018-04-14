package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println(v)
	}
}

func Server(queue chan int) {
	var sem = make(chan int, 4)
	for i := range queue {
		sem <- 1
		go func(i int) {
			fmt.Println(i)
			time.Sleep(1e9)
			<-sem
		}(i)
	}
}

func main() {
	//Channel作为goroutine间的一种通信机制，与操作系统的其它通信机制类似，一般有两个目的：同步，或者传递消息。
	var ch chan int
	ch = make(chan int, 1)
	//var rch <-chan int
	//rch = make(chan int)
	//var wch chan<- int
	//wch = make(chan int, 1)
	ch <- 1
	close(ch)
	//rch <- 0 //error
	fmt.Println(<-ch)
	//fmt.Println(<-rch)
	//wch <- 0
	//close(wch)

	//fmt.Println(<-wch)

	//var queue chan int = make(chan int, 10)
	//go Producer(queue)
	//go Consumer(queue)
	//time.Sleep(1e11)

	//var t = make([]int, 0, 10)
	//var s = make([]int, 0, 10)
	//fmt.Printf("addr:%p \t\tlen:%v content:%v\n", t, len(t), t);
	//fmt.Printf("addr:%p \t\tlen:%v content:%v\n", s, len(s), s);
	//t = append(s, 1, 2, 3, 4)
	//fmt.Println(t)
	//fmt.Println(s)
	//fmt.Printf("addr:%p \t\tlen:%v content:%v\n", t, len(t), t);
	//fmt.Printf("addr:%p \t\tlen:%v content:%v\n", s, len(s), s);

	//var queue = make(chan int)
	//go func() {
	//	i := 1
	//	for {
	//		queue <- i
	//		i++
	//	}
	//}()
	//go Server(queue)
	i := 0
	i1 := make(chan int)
	//i2 := make(chan int)

	go func() {
		for {
			<-i1
			i++
			fmt.Println("1", i)
			//time.Sleep(1e9)
			i1 <- 1
		}

	}()
	go func() {
		for {
			<-i1
			i++
			fmt.Println("2", i)
			//time.Sleep(1e9)
			i1 <- 1
		}
	}()
	go func() {
		for {
			<-i1
			i++
			fmt.Println("3", i)
			//time.Sleep(1e9)
			i1 <- 1
		}
	}()
	//go func() {
	//	for {
	//		<-i2
	//		i++
	//		fmt.Println("2", i)
	//		time.Sleep(1e9)
	//		i1 <- 1
	//	}
	//}()

	i1 <- 1

	time.Sleep(1e11)
}
