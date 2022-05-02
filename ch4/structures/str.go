package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

var s1 aStructure

func main() {
	p1 := aStructure{"fmt", 12, -2}
	fmt.Println(s1.height)
	fmt.Println(p1)
}