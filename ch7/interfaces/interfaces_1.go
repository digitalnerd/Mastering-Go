package main

import (
	"fmt"
	"time"
)

type Printer interface {
	Print()
}

type User struct {
	name string
	lastName string
	age int
}

type Document struct {
	name string
	documentType string
	date time.Time
}

func (d Document) Print() {
	fmt.Printf("Document name: %s, type: %s, date: %s \n", d.name, d.documentType, d.date)
}

func (u User) Print() {
	fmt.Printf("Hi I am %s %s and I am %d years old \n", u.lastName, u.name, u.age)
}

func Process(obj Printer) {
	obj.Print()
}

func main() {
	u := User{name: "John", age: 24, lastName: "Smith"}
  	doc := Document{name: "doc.csv", documentType: "csv", date: time.Now()}
  	Process(u)
  	Process(doc)
}