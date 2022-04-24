package main

import "fmt"

type twoInt struct {
	X int64
	Y int64
}

func regularFunction(a, b twoInt) twoInt {
	temp := twoInt{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

// The method() function is equivalent to the regularFunction()
// However, the method() function is a type method
func (a twoInt) method(b twoInt) twoInt {
	temp := twoInt{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func main() {
	i := twoInt{X: 1, Y: 2}
	j := twoInt{X: -5, Y: -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}