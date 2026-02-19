// interfaces/type-switch/begin/main.go
package main

import "fmt"
// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
    func whatAmI(i interface{}) string{
			switch i.(type) {
			case int:
				return fmt.Sprintf("I am an int and my value is %d", i)
			case string:
				return fmt.Sprintf("I am a string and my value is %s", i)
			default:
				return "I am of a different type"
			}

		}
func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	 fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
