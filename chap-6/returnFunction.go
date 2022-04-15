package main

import (
	"fmt"
)

// As you can see from the implementation of funReturnFun(), its return value is an
// anonymous function (func() int).
func funcReturnFunc() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {

	i := funcReturnFunc()
	j := funcReturnFunc()

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())

}
