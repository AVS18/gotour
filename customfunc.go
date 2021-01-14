package main
import (
	"fmt"
)
func adder(x int, y int) int{
	return (x+y)
}
func main(){
	fmt.Println("Addtion of 2,3 is ",adder(2,3))
}