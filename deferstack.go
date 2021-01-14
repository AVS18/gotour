package main
import "fmt"
func main(){
	for i := 10;i>0;i-- {
		defer fmt.Println(i)
	}
}