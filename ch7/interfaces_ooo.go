package main

import "fmt"

// defining an interface
type Sport interface {
	// name of sport method
	sportName() string
}

type Human struct {
	name string
	sport string
}

func (h Human) sportName() string {
	return h.name + " plays " + h.sport + "."
}

func main() {
	// declaring a struct instance
	human1 := Human{"Rahul", "chess"}
	fmt.Println(human1.sportName())

	// declaring another struct instance
    human2 := Human{"Riya", "carrom"}
    // printing details of human2
    fmt.Println(human2.sportName())
}