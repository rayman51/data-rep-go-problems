//https://stackoverflow.com/questions/46395819/get-sum-of-bigint-number-golang
package main

import (
	"fmt"
	"math/big"
)
/*func main() {
	var num int
	var fact int
	fmt.Println("Enter the number to factorialise ...") // prompts user
	fmt.Scanf("%d\n", &num)

	if num == 0 {// if num = 0 prompts user to enter again 
	
		fmt.Println("   Invalid Entry Please Enter again : ")
		fmt.Scanf("%d", &num);
		fmt.Println("====================================\n ")
	}// if

	if num>0 {// if valid number run the while loop 
	
		// loop to multiply numbers
		for num > 0{
			fact *= num//multiplies number 
			num--// decrements number
		}// while
	}// if
	fmt.Printf("    Factorial is : %d", fact)

} // main*/

func factoral(n uint64) (r *big.Int) {

	one, bn := big.NewInt(1), new(big.Int).SetUint64(n)

	r = big.NewInt(1)
	if bn.Cmp(one) <= 0 {
		return
	}
	for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i, one) {
		r.Mul(r, i)
	}
	return
}

func add(number *big.Int) *big.Int {
	ten := big.NewInt(10)
	sum := big.NewInt(0)
	mod := big.NewInt(0)
	for ten.Cmp(number) < 0 {
		sum.Add(sum, mod.Mod(number, ten))
		number.Div(number, ten)
	}
	sum.Add(sum, number)
	return sum
}
func main() {
	fmt.Println("The factorial of 100 is ", factoral(100))

	fmt.Println("The sum of factorial 100 is ", add(factoral(100)))

}
