package main
import "fmt"
func main(){
	s := []int{0,1,2,3}
	s = append(s,4)
	fmt.Println(s)
	s = append(s,5,6)
	fmt.Println(s)
	s = append(s,7,8,9,10)
	fmt.Println(s)
}