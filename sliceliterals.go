package main
import "fmt"
func main(){
	a := []int{1,2,3,4,5}
	b := []bool{false,true,false,false,false}
	s := []struct{
		i int
		b bool
	}{
		{1,false},
		{2,true},
		{3,false},
	}
	fmt.Println("Struct",s)
	fmt.Println(a,b)
	
}