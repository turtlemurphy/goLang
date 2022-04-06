package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	fmt.Print("\n")
	fmt.Print("\n")

	for j := 1; j <= 10; j++ {

		fmt.Print(j, " ")

		switch j {
		case 1:
			fmt.Println(" one")
		case 2:
			fmt.Println(" two")
		case 3:
			fmt.Println(" three")
		case 4:
			fmt.Println(" four")
		case 5:
			fmt.Println(" five")
		case 6:
			fmt.Println(" six")
		case 7:
			fmt.Println(" seven")
		case 8:
			fmt.Println(" eight")
		case 9:
			fmt.Println(" nine")
		case 10:
			fmt.Println("ten")
		}

		if j%20 == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")
	fmt.Print("\n")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println(time.Now())
		fmt.Println("It's the weekend")
	default:
		fmt.Println(time.Now())
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Print(t, " is a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
