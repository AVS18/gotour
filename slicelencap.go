package main
import "fmt"
func main(){
	a := []int{0,1,2,3,4}
	print(a[:4])
	print(a[1:])
	print(a[1:4])	
}
func print(s []int) {
	fmt.Printf("Capacity %v, Length %v, Array %v\n",cap(s),len(s),s);
}