package main
import "fmt"
func main(){
	var array [4]int;
	for i:=0 ; i<4;i++{
		array[i]=i;
	}
	//slice aray
	fmt.Println(array[1:3])
}