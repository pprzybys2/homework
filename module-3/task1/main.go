package main

import "fmt"

func main() {
	fmt.Println("TEST")
	x1 := 5

	x2 := &x1

	x3 := &x1

	fmt.Println(x1, x2, x3)

	fmt.Println(&x1, &x2, &x3)

	fmt.Println(x1, *x2, *x3)
}

func increment(x int) {
	x++

	fmt.Println(&x, x)
}
