package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title:", b.title)
	fmt.Println("Content:", b.content)
	fmt.Println("Author:", b.fullName())
	fmt.Println("Bio:", b.bio)
}

type website struct {
	blogPost []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.blogPost {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"K",
		"L",
		"Golang developer",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		blogPost: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
	//blogPost1.details()
}
