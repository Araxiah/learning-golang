package main

import (
	"fmt"
	"time"
)
func main() {
	i := 2
	fmt.Println("write ", 2, " as")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend my dude!")
	default:
		fmt.Println("It's the weekday c:")
	}

	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("it's afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Println("I don't know what the type is!! ahhhhh!!!\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}