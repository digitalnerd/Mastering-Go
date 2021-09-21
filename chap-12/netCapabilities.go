package main

import (
	"fmt"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags:", i.Name)
		fmt.Println("Interface MTU:", i.Flags.String())
		fmt.Println("Interface Hardware Address:", i.HardwareAddr)

		fmt.Println()
	}

}
