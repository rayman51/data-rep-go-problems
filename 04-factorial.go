package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := new(big.Int)
	x.MulRange(1, 100)
	fmt.Println(x)

}
