package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

// emulate a typical while loop
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}

		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

// do .. while loop
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

// range keyword
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index: ", i, "value: ", value)
	}
}
