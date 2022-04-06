package main

import (
	"fmt"
	"time"
)

func main() {
	j := 2
	fmt.Print("Write ", j, " as ")
	fmt.Print("\n")
	fmt.Print("\n")

	for j := 1; j <= 10; j++ {

		fmt.Print(j, " ")

		switch j {
		case 1:
			fmt.Print(" one")
			if j%2 == 0 {
				fmt.Print("   is even")
				fmt.Print("\n")
			} else {
				fmt.Print("   is odd")
				fmt.Print("\n")
			}
		case 2:
			fmt.Print(" two")
			if j%2 == 0 {
				fmt.Print("   is even")
				fmt.Print("\n")
			} else {
				fmt.Print("   is odd")
				fmt.Print("\n")
			}
		case 3:
			fmt.Print(" three")
			if j%2 == 0 {
				fmt.Print(" is even")
				fmt.Print("\n")
			} else {
				fmt.Print(" is odd")
				fmt.Print("\n")
			}
		case 4:
			fmt.Print(" four")
			if j%2 == 0 {
				fmt.Print("  is even")
				fmt.Print("\n")
			} else {
				fmt.Print("  is odd")
				fmt.Print("\n")
			}
		case 5:
			fmt.Print(" five")
			if j%2 == 0 {
				fmt.Print("  is even")
				fmt.Print("\n")
			} else {
				fmt.Print("  is odd")
				fmt.Print("\n")
			}
		case 6:
			fmt.Print(" six")
			if j%2 == 0 {
				fmt.Print("   is even")
				fmt.Print("\n")
			} else {
				fmt.Print("   is odd")
				fmt.Print("\n")
			}
		case 7:
			fmt.Print(" seven")
			if j%2 == 0 {
				fmt.Print(" is even")
				fmt.Print("\n")
			} else {
				fmt.Print(" is odd")
				fmt.Print("\n")
			}
		case 8:
			fmt.Print(" eight")
			if j%2 == 0 {
				fmt.Print(" is even")
				fmt.Print("\n")
			} else {
				fmt.Print(" is odd")
				fmt.Print("\n")
			}
		case 9:
			fmt.Print(" nine")
			if j%2 == 0 {
				fmt.Print("  is even")
				fmt.Print("\n")
			} else {
				fmt.Print("  is odd")
				fmt.Print("\n")
			}
		case 10:
			fmt.Print("ten")
			if j%2 == 0 {
				fmt.Print("   is even")
				fmt.Print("\n")
			} else {
				fmt.Print("   is odd")
				fmt.Print("\n")
			}
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

	whatAmI := func(j interface{}) {
		switch t := j.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println(t, " is a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("supercalifragilisticexpialidotious")
}
