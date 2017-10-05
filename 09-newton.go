//newton’s method for calculating square roots
//Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.
//zNext = z - ((z*z - x) / (2 * z))

package main

import (
	"fmt"
	"math"
)

func zNext(z float64, x float64) float64 { //z float64, x float64...returns float64

	return z - ((z*z - x) / (2 * z))
}

func main() {


	var x float64
	var z float64
	var guess float64


	//scan user input for square root
	fmt.Printf("Please enter the number to find square root of: ")
	fmt.Scanf("%f", &x)

	//scan user input for initial guess
	fmt.Printf("Please enter your guess: ")
	fmt.Scan(&guess)
	fmt.Println()
	//fmt.Scanf("%f", &startingGuess) //doesn't work

	for z = guess; z != zNext(z, x); z = zNext(z, x) {

		//print the guess on each iteration
		fmt.Printf("Current guess: %13f\n", z)

	}
	//finally, z is a good appriximation of th esquare root of x
	fmt.Printf("The square root of %f is %f\n", x, z)

	//print out the math.Sqrt value
	fmt.Printf("math.Sqrt gives the value of square root %f as %f\n", x, math.Sqrt(x))

}
