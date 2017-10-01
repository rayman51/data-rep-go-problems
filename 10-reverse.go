//https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
package main

import "fmt" // imports package to print to screen

func main() {
	fmt.Println("====================================")
	fmt.Println("          WORD REVERESER")
	fmt.Println("====================================")
	fmt.Println("Enter a word you wish to reverse :") // prompts user
	var input string
	fmt.Scanf("%s\n", &input)

	n := 0
	rune := make([]rune, len(input)) // makes array to store the word
	for _, r := range input {        // runs through the word
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i] // reads back through the array
	}
	// Convert back to UTF-8.
	output := string(rune) // outputs the word in reverse
	fmt.Println("====================================")
	fmt.Printf("The reverse of %s is %s \n", input, output)
	fmt.Println("====================================")
} // main
