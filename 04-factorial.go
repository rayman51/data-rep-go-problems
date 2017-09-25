//https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion

package main

import (
	"fmt"      // imports fmt package
	"math/big" // imports mathbig package
)

func main() {
	factorial()
} // main
func factorial() {
	x := new(big.Int)
	x.MulRange(1, 100)
	fmt.Println(x)

} // factorial
