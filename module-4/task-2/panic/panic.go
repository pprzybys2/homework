package panic

import "fmt"

func iWillPanic() {
	panic("something")
}

func catchPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()
	iWillPanic()
}
