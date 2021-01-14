package main
import "fmt"
type node struct{
	data int
	float_data float64
}
func main(){
	var t = node{1,1.3}
	fmt.Println(t.data,t.float_data)
}