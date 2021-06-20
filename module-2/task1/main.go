package main

import (
	"computator/fibonacci"
	"fmt"
	"strconv"
)

func main() {
	var in string
	var num interface{}
	fmt.Println("What number?")
	fmt.Scanf("%s", &in)
	if n, err := strconv.ParseInt(in, 10, 64); err == nil {
		num = int(n)
	} else if n, err := strconv.ParseUint(in, 10, 64); err == nil {
		num = uint(n)
	} else if n, err := strconv.ParseFloat(in, 64); err == nil {
		num = n
	} else {
		fmt.Println("Couldn't parse number")
	}
	out := fibonacci.Fibonacci(num)
	fmt.Printf("Number is %v", out)

}
