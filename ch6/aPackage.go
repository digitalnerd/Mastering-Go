package aPackage

import "fmt"

func A() {
	fmt.Println("This is a function A!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21
