package fibonacci

import "fmt"

func Fibonacci(n interface{}) interface{} {
	switch t := n.(type) {
	case int:
		//fmt.Printf("CO TO %v", t)
		return fibonnaciInt(t)
	case uint:
		return fibonnaciUint(t)
	case float64:
		return fibonnaciFloat(t)
	default:
		return fmt.Errorf("Unknow type")
	}
}

func fibonnaciInt(n int) int {
	//fmt.Printf("INTTT %v", n)
	if n <= 1 {
		return n
	} else {
		return fibonnaciInt(n-1) + fibonnaciInt(n-2)
	}
}

func fibonnaciUint(n uint) uint {
	if n <= 1 {
		return n
	} else {
		return fibonnaciUint(n-1) + fibonnaciUint(n-2)
	}
}

func fibonnaciFloat(n float64) float64 {
	if n <= 1 {
		return n
	} else {
		return fibonnaciFloat(n-1) + fibonnaciFloat(n-2)
	}
}
