package main

import (
	. "fmt"
)

func main() {
	i := 0
	for ; i < 6; i++ {
		defer func() {
			Println(i)
		}()
	}
}
