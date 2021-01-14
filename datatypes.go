package main
import (
	"fmt"
	"math/cmplx"
)

var (
	tobe   bool       = false
	maximum_int uint64     = 1<<64 - 1
	a      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", tobe,tobe)
	fmt.Printf("Type: %T Value: %v\n", maximum_int,maximum_int)
	fmt.Printf("Type: %T Value: %v\n", a,a)
}
