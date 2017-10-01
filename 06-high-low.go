//https://gist.github.com/pyk/10519339
package main

import "fmt" // imports fmt

func main() {
	var n, smallest, biggest int //variables used
	x := []int{                  // array of number to search through
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for _, v := range x { // runs through the array and looks for the largest number
		if v > n {
			fmt.Println(v, ">", n) // if the present number is larger than the previous
			n = v                  //the new number become the largest for now
			biggest = n
		} else {
			fmt.Println(v, "<", n)
		}
	}

	fmt.Println("The biggest number is ", biggest) // prints lagest number to screen

	for _, v := range x { // runs through the array and looks for the smallest number
		if v > n {
			fmt.Println(v, ">", n) // if the present number is smaller than the previous
		} else { //the new number become the smallest for now
			fmt.Println(v, "<", n)
			n = v
			smallest = n
		}
	}
	fmt.Println("The smallest number is ", smallest) // prints lagest number to screen
} // main
