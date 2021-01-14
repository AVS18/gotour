package main
import (
	"fmt"
	"math"
)
func main() {
	var a, b int = 13, 5
	var c float64 = math.Sqrt(float64(a*a + b*b))
	var d uint = uint(c)
	fmt.Printf("Distance between %v and %v is %v",a,b,d)
}
