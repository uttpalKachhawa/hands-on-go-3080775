// types/structs/fields/begin/main.go
package main

import "fmt"
// define a struct type for author
type author struct {
	name string
	age int
}

func main() {
	// intialize author
	a := author{
		name: "uttpal",
		age:  25,
	}

	// print the author
	fmt.Printf("%#v\n %#v\n", a.name, a.age)
	fmt.Printf("%#v\n", a)
}
