//http://www.golangpro.com/2016/01/check-string-palindrome-go.html?m=1
package main

import (
	"fmt"     // imports fmt
	"strings" // imports strings
)

func main() {

	var ip string // variables used
	fmt.Println("=======================")
	fmt.Println("  PALINDROME CHECKER")
	fmt.Println("=======================")
	fmt.Println("Enter the word :") // user prompt
	fmt.Scanf("%s\n", &ip)
	ip = strings.ToLower(ip)
	fmt.Println(isP(ip)) // passes ip to the function isP
}

//Function to test if the string entered is a Palindrome

func isP(s string) string {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ { //checks if the last letter and the first letter are the same
		if s[i] != s[last-i] { // then moves on and check2nd letter and 2nd last letter and ect...
			return "NO. It's not a Palimdrome."
		}
	}
	return "YES! You've entered a Palindrome"
}
