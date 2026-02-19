// interfaces/basics/begin/main.go
package main

import "fmt"
// define a humanoid interface with speak and walk methods returning string
type humanoid interface{
	speak()
	walk()
}
// define a person type that implements humanoid interface
type person struct{
	name string
}

//way of implementing the humanoid interface for person type
func (p person) speak() {
	fmt.Printf("%s speaking....", p.name)
}

func(p person) walk() {
fmt.Printf("%s walking....", p.name)
}
// implement the Stringer interface for the person type

// define a dog type that can walk but not speak
type dog struct{
	name string
}
func(d dog) walk() {
	fmt.Printf("%s walking....", d.name)
}

func(d dog) speak() {
	fmt.Printf("%s speaking....", d.name)
}
func main() {
	// invoke with a person
	p:= person{name: "John"}
	doHumanThings(p)

	d:= dog{name: "Rex"}
	// can we invoke with a dog?
	doHumanThings(d)

	// fmt.Println(p)
}

func doHumanThings(h humanoid) {
 	h.speak()
	h.walk()
 }
