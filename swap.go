package main
import "fmt"

func swap(a string, b string)(string,string){
	return b, a
}
func main(){
	var a string = "aditya"
	var b string = "2"
	fmt.Println("Before swap ",a,b)
	a,b = swap(a, b)
	fmt.Println("After swap ",a,b)
}