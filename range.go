package main
import "fmt"
func main(){
	a := [5]int{0,1,2,3,4}
	for i := range a {
		fmt.Println(i,a[i])
	}
}