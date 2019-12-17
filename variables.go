package main

import "fmt"

func main() {
	var a = "first"
	fmt.Println("I am var a-",a)

	var b, c = "second", "third"
	fmt.Println("I am var b-",b)
	fmt.Println("I am var c-",c)
	fmt.Println("I am vars b then c-",b, c)

	var d int = 4
	fmt.Println("I am var d stated as int",d)

	e := "apple"
	fmt.Println("I am var e writen with a shorthand syntax(:=)", e)

	var f int
	fmt.Println("I am var f-",f)

}