package main

import (
	"fmt"
)

func returnPtr(x int) *int {
	y := x * x
	return &y
}

func main() {
	sq := returnPtr(10)
	// The * character dereferences a pointer variable, which means that it returns the actual value
	// stored at the memory address instead of the memory address itself.
	fmt.Println("sq value:", *sq)
	// Return the memory address of the sq variable, not the int value stored in it.
	fmt.Println("sq memory address:", sq)
}
