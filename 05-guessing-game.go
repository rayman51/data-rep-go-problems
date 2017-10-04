//http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html
package main

import (
	"fmt"       // imports
	"math/rand" // packages
	"time"      // needed
) // import

func xrand(min, max int) int { //function to
	rand.Seed(time.Now().Unix()) //create a random number
	return rand.Intn(max-min) + min
} //xrand

func main() {
	myrand := xrand(1, 10) // call function and give it the range of 1-10

	tries := 0    // variables
	var guess int //  used
	var prev int 
	fmt.Println("====================================")
	fmt.Println("Welcome to My Guess The Number Game!")
	fmt.Println("====================================")

	for guess != myrand {
		fmt.Println("Take a guess...") // prompts user
		fmt.Scanf("%d\n", &guess)
		
		if guess == prev{
			fmt.Println("You already tried this number\n ")
			tries--
		}
		if guess > myrand { // checks if guess is higher than random number
			fmt.Println("Too high")
		} else if guess < myrand { // checks if guess is lower than random number
			fmt.Println("Too low")
		} else { // when number is picked, informs the user and breaks the loop
			fmt.Println("====================================")
			fmt.Printf("The number was %d \n", myrand)
			fmt.Printf("Good job! You guessed it in %d tries", tries)
			break
		} //end ofif/else
		tries++ //counts tries
		prev = guess
	} // end of for loop

} // main
