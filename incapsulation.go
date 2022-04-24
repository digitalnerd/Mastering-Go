package main

import "fmt"

type MyObject struct {
	MyStr string
	MyInt int
}

func main() {
	// myObj := new(MyObject)
	// fmt.Println(myObj)
	// myObj.MyStr = "Go is super!"
	// myObj.MyInt = 11
	// fmt.Println(myObj)

	myObj := MyObject{
		MyStr: "Go is super!",
	}
	fmt.Println(&myObj)
}