//http://wiki.c2.com/?FizzBuzzTest
package main

import "fmt" // import package to print to screen

func main() {
	i := 1         // initialises i at 1
	for i <= 100 { // runs through a loop until i = 100
		if i%3 == 0 && i%5 == 0 { //triggers if number is divisable by 3 & 5 with no remainder
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 { //triggers if number is divisable by 3 with no remainder
			fmt.Println("Fizz")
		} else if i%5 == 0 { //triggers if number is divisable 5 with no remainder
			fmt.Println("Buzz")
		} else {
			fmt.Println(i) //prints numbers to screen
		} // if/else
		i++ /// increments count
	} // for
} // main
