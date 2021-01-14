package main
import "fmt"

func namedReturn(sum int)(x , y int){
	x = sum - 10
        y = sum * 4 - 2
	return
}
func main(){
	fmt.Println("AVS Aditya - Named Return")
	fmt.Println(namedReturn(100))
}