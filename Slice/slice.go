package main

import (
	. "fmt"
)

func main() {
	s := []int{5}
	Println("cap(s) =", cap(s), "ptr(s) =", &s[0], s, "ptr(x) =")

	s = append(s, 7)
	Println("cap(s) =", cap(s), "ptr(s) =", &s[0], s, "ptr(x) =")

	s = append(s, 9)
	Println("cap(s) =", cap(s), "ptr(s) =", &s[0], s, "ptr(x) =")

	x := append(s, 11)
	Println("cap(s) =", cap(s), "ptr(s) =", &s[0], s, "ptr(x) =", &x[0], x)

	y := append(s, 12)
	Println("cap(s) =", cap(s), "ptr(s) =", &s[0], s, "ptr(x) =", &x[0], x, "ptr(x) =", &y[0], y)

	t := []int {1,2,3}
	Println("cap(s) =", cap(t), "ptr(s) =", &t[0], t)

	t = append(t,2)
	Println("cap(s) =", cap(t), "ptr(s) =", &t[0], t)
	t = append(t,3)
	Println("cap(s) =", cap(t), "ptr(s) =", &t[0], t)

	m := append(t,4)
	Println("cap(s) =", cap(t), "ptr(s) =", &t[0], t, "ptr(x) =", &m[0], m, "ptr(x) =")

	n := append(t,5)
	Println("cap(s) =", cap(t), "ptr(s) =", &t[0], t, "ptr(x) =", &m[0], m, "ptr(x) =", &n[0], n)
}
